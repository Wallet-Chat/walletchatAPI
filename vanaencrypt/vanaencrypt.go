package vanaencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateSharedSecret computes the ECDH shared secret between a private key and a public key.
func GenerateSharedSecret(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey.Curve != publicKey.Curve {
		return nil, errors.New("private and public keys must use the same curve")
	}

	// Perform scalar multiplication
	x, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())

	// Return the X coordinate as the shared secret
	return x.Bytes(), nil
}

// DerivePadded computes a shared secret between a private key and a public key, with padding.
func DerivePadded(privateKeyHex string, publicKeyHex string) ([]byte, error) {
	// Decode private key
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, errors.New("invalid private key hex")
	}
	if len(privateKeyBytes) != 32 {
		return nil, errors.New("private key must be 32 bytes")
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	// Decode public key
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return nil, errors.New("invalid public key hex")
	}
	if len(publicKeyBytes) != 65 && len(publicKeyBytes) != 33 {
		return nil, errors.New("public key must be 65 (uncompressed) or 33 (compressed) bytes")
	}

	var publicKey *ecdsa.PublicKey
	// If the public key is compressed, decompress it
	if len(publicKeyBytes) == 33 {
		publicKey, err = crypto.DecompressPubkey(publicKeyBytes)
		if err != nil {
			return nil, errors.New("failed to decompress public key")
		}
	} else {
		publicKey, err = crypto.UnmarshalPubkey(publicKeyBytes)
		if err != nil {
			return nil, errors.New("invalid uncompressed public key")
		}
	}

	// Derive shared secret
	x, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())

	// Convert the shared secret to a padded 32-byte buffer
	sharedSecret := x.Bytes()
	paddedSecret := make([]byte, 32)
	copy(paddedSecret[32-len(sharedSecret):], sharedSecret)

	return paddedSecret, nil
}

// EncryptWithPubKey performs ECIES encryption using a public key.
func EncryptWithPubKey(publicKeyTo []byte, msg []byte, opts map[string][]byte) (map[string][]byte, error) {
	if len(publicKeyTo) == 0 || len(msg) == 0 {
		return nil, errors.New("public key and message are required")
	}

	// Load the public key
	publicKey, err := crypto.UnmarshalPubkey(publicKeyTo)
	if err != nil {
		return nil, err
	}

	// Generate an ephemeral private key
	ephemPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	ephemPublicKey := crypto.FromECDSAPub(&ephemPrivateKey.PublicKey)

	// Derive shared secret
	sharedSecret, err := GenerateSharedSecret(ephemPrivateKey, publicKey)
	if err != nil {
		return nil, err
	}

	// Hash the shared secret to derive encryption and MAC keys
	hash := sha512.Sum512(sharedSecret)
	encryptionKey := hash[:32]
	macKey := hash[32:]

	// Generate IV
	iv := opts["iv"]
	if len(iv) == 0 {
		iv = make([]byte, aes.BlockSize)
		if _, err := rand.Read(iv); err != nil {
			return nil, err
		}
	}

	// Encrypt the message using AES-CBC
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	paddedMsg := pad(msg, aes.BlockSize)
	ciphertext := make([]byte, len(paddedMsg))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedMsg)

	// Prepare data for MAC
	dataToMac := bytes.Join([][]byte{iv, ephemPublicKey, ciphertext}, []byte{})

	// Compute HMAC
	h := hmac.New(sha256.New, macKey)
	h.Write(dataToMac)
	mac := h.Sum(nil)

	// Return the ECIES components
	return map[string][]byte{
		"iv":             iv,
		"ephemPublicKey": ephemPublicKey,
		"ciphertext":     ciphertext,
		"mac":            mac,
	}, nil
}

// EncryptWithWalletPublicKey encrypts data using a wallet public key and returns the encrypted data as a hex string.
func EncryptWithWalletPublicKey(data string, publicKeyHex string) (map[string][]byte, error) {
	empty := map[string][]byte{}

	// Remove "0x" prefix if present
	if len(publicKeyHex) > 2 && publicKeyHex[:2] == "0x" {
		publicKeyHex = publicKeyHex[2:]
	}

	if len(publicKeyHex) == 0 {
		return empty, errors.New("public key is required")
	}

	// Decode the public key from hex
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return empty, err
	}

	// Ensure the public key is in uncompressed format
	if len(publicKeyBytes) == 64 {
		publicKeyBytes = append([]byte{0x04}, publicKeyBytes...)
	}

	// Convert data to bytes
	messageBytes := []byte(data)

	// Call the existing EncryptWithPubKey function
	encryptedData, err := EncryptWithPubKey(publicKeyBytes, messageBytes, nil)
	if err != nil {
		return empty, err
	}

	// Concatenate IV, ephemeral public key, ciphertext, and MAC
	return encryptedData, nil
}

// ClientSideEncrypt encrypts the input bytes using a password and returns encrypted bytes.
func ClientSideEncrypt(inputBytes []byte, password string) ([]byte, error) {
	if len(inputBytes) == 0 {
		return nil, errors.New("input bytes cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// Create OpenPGP message
	config := &packet.Config{}
	messageBuffer := new(bytes.Buffer)
	messageWriter, err := openpgp.SymmetricallyEncrypt(messageBuffer, []byte(password), nil, config)
	if err != nil {
		return nil, err
	}

	_, err = messageWriter.Write(inputBytes)
	if err != nil {
		return nil, err
	}

	err = messageWriter.Close()
	if err != nil {
		return nil, err
	}

	// Return encrypted bytes
	return messageBuffer.Bytes(), nil
}

// pad adds padding to the message to make its length a multiple of blockSize.
func pad(msg []byte, blockSize int) []byte {
	padding := blockSize - len(msg)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(msg, padtext...)
}
