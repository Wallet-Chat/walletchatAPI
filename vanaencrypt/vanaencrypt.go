package vanaencrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	protonPgpCrypto "github.com/ProtonMail/gopenpgp/v3/crypto"
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
func EncryptWithPubKey(publicKeyTo []byte, msg []byte, opts map[string][]byte, staticIV []byte, ephemPrivateKeyBytes []byte) (map[string][]byte, error) {
	if len(publicKeyTo) == 0 || len(msg) == 0 {
		return nil, errors.New("public key and message are required")
	}

	// Load the public key
	publicKey, err := crypto.UnmarshalPubkey(publicKeyTo)
	if err != nil {
		return nil, err
	}

	// Convert the 32-byte ephemeral private key to ECDSA private key
	ephemPrivateKey, err := crypto.ToECDSA(ephemPrivateKeyBytes)
	if err != nil {
		return nil, errors.New("invalid ephemeral private key")
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

	// Use the provided static IV
	iv := staticIV // Use the static IV passed as a parameter

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
func EncryptWithWalletPublicKey(data string, publicKeyHex string, iv []byte, ephemPrivateKeyBytes []byte) (map[string][]byte, []byte, error) {
	empty := map[string][]byte{}

	// Remove "0x" prefix if present
	if len(publicKeyHex) > 2 && publicKeyHex[:2] == "0x" {
		publicKeyHex = publicKeyHex[2:]
	}

	// Decode the public key from hex
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return empty, nil, err
	}

	// Ensure the public key is in uncompressed format
	if len(publicKeyBytes) == 64 {
		publicKeyBytes = append([]byte{0x04}, publicKeyBytes...)
	}

	// Convert data to bytes
	messageBytes := []byte(data)

	// Initialize the IV and ephemeral private key with random bytes if was not provided
	if len(iv) != 16 {
		fmt.Println("Creating Random IV")
		iv = make([]byte, 16) // 16 bytes for IV
		if _, err := rand.Read(iv); err != nil {
			return empty, nil, err
		}
	}

	if len(ephemPrivateKeyBytes) != 32 {
		fmt.Println("Creating Random ephemeralPrivKey")
		ephemPrivateKeyBytes = make([]byte, 32) // 32 bytes for ephemeral private key
		if _, err := rand.Read(ephemPrivateKeyBytes); err != nil {
			return empty, nil, err
		}
	}

	//for test use known values to compare with web/example version
	// iv = []byte{
	// 	169, 138, 29, 49, 139, 11, 183, 51,
	// 	167, 5, 144, 163, 203, 214, 217, 224,
	// }
	// ephemPrivateKeyBytes = []byte{
	// 	147, 207, 81, 186, 169, 91, 245, 42,
	// 	148, 220, 122, 136, 222, 82, 10, 86,
	// 	230, 210, 241, 85, 15, 154, 77, 60,
	// 	38, 91, 211, 211, 243, 2, 214, 203,
	// }
	//end hardcoded test values

	// Call the existing EncryptWithPubKey function with the static IV and ephemeral private key
	encryptedData, err := EncryptWithPubKey(publicKeyBytes, messageBytes, nil, iv, ephemPrivateKeyBytes)
	if err != nil {
		return empty, nil, err
	}

	// Convert encrypted data to hex string and print
	// for key, value := range encryptedData {
	// 	fmt.Printf("%s: %s\n", key, hex.EncodeToString(value))
	// }

	// Return the encrypted data
	return encryptedData, ephemPrivateKeyBytes, nil
}

func DecryptWithPrivKey(privateKeyBytes []byte, encryptedData map[string][]byte) ([]byte, error) {
	// Validate input
	if len(privateKeyBytes) == 0 {
		return nil, errors.New("private key is required")
	}
	if encryptedData == nil {
		return nil, errors.New("encrypted data is required")
	}

	// Extract encrypted components
	iv := encryptedData["iv"]
	ephemPublicKeyBytes := encryptedData["ephemPublicKey"]
	ciphertext := encryptedData["ciphertext"]
	mac := encryptedData["mac"]

	// Ensure all components are present
	if iv == nil || ephemPublicKeyBytes == nil || ciphertext == nil || mac == nil {
		return nil, errors.New("incomplete encrypted data")
	}

	// Convert the 32-byte private key to ECDSA private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	// Load the ephemeral public key
	ephemPublicKey, err := crypto.UnmarshalPubkey(ephemPublicKeyBytes)
	if err != nil {
		return nil, errors.New("invalid ephemeral public key")
	}

	// Derive the shared secret
	sharedSecret, err := GenerateSharedSecret(privateKey, ephemPublicKey)
	if err != nil {
		return nil, err
	}

	// Hash the shared secret to derive encryption and MAC keys
	hash := sha512.Sum512(sharedSecret)
	encryptionKey := hash[:32]
	macKey := hash[32:]

	// Verify MAC for integrity
	dataToMac := bytes.Join([][]byte{iv, ephemPublicKeyBytes, ciphertext}, []byte{})
	h := hmac.New(sha256.New, macKey)
	h.Write(dataToMac)
	calculatedMac := h.Sum(nil)
	if !hmac.Equal(calculatedMac, mac) {
		return nil, errors.New("MAC verification failed")
	}

	// Decrypt the ciphertext using AES-CBC
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintextPadded := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintextPadded, ciphertext)

	// Unpad the plaintext
	plaintext, err := unpad(plaintextPadded, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Helper function for PKCS#7 padding removal
func unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	if len(data)%blockSize != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}

	paddingLen := int(data[len(data)-1])
	if paddingLen > blockSize || paddingLen == 0 {
		return nil, errors.New("invalid padding length")
	}

	for _, v := range data[len(data)-paddingLen:] {
		if int(v) != paddingLen {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:len(data)-paddingLen], nil
}

func ClientSideEncrypt(data []byte, password string) ([]byte, error) {
	// Create a password-protected keyring
	passphrase := []byte(password)
	pgp := protonPgpCrypto.PGP()
	// Encrypt data with a password
	encHandle, err := pgp.Encryption().Password(passphrase).New()
	if err != nil {
		fmt.Println("error during passphrase init!!!!!!!!!!!!!!!!!!")
	}
	pgpMessage, err := encHandle.Encrypt(data)
	if err != nil {
		fmt.Println("error during encryption!!!!!!!!!!!!!!!!!!")
	}
	armored, err := pgpMessage.ArmorBytes()

	return armored, err
}

func ClientSideDecrypt(data []byte, password string) ([]byte, error) {
	// Create a password-protected keyring
	passphrase := []byte(password)
	pgp := protonPgpCrypto.PGP()

	// Decrypt data with a password
	decHandle, err := pgp.Decryption().Password(passphrase).New()
	decrypted, err := decHandle.Decrypt(data, protonPgpCrypto.Armor)
	myMessage := decrypted.Bytes()

	return myMessage, err
}

// pad adds padding to the message to make its length a multiple of blockSize.
func pad(msg []byte, blockSize int) []byte {
	padding := blockSize - len(msg)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(msg, padtext...)
}

// EncryptWithPGPPublicKey encrypts the given data using the provided PGP public key.
func EncryptWithPGPPublicKey(publicKeyPEM string, data []byte) (string, error) {
	// Load the public key from the PEM string
	publicKey, err := protonPgpCrypto.NewKeyFromArmored(publicKeyPEM)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		fmt.Println("Public Key Content:\n", publicKeyPEM)
		return "", err
	}

	pgp := protonPgpCrypto.PGP()
	// Encrypt plaintext message using a public key
	encHandle, err := pgp.Encryption().Recipient(publicKey).New()
	if err != nil {
		fmt.Println("error during EncryptWithPGPPublicKey 2")
		return "", err
	}
	pgpMessage, err := encHandle.Encrypt(data)
	if err != nil {
		fmt.Println("error during EncryptWithPGPPublicKey 3")
		return "", err
	}
	armored, err := pgpMessage.ArmorBytes()
	if err != nil {
		fmt.Println("error during EncryptWithPGPPublicKey 4")
		return "", err
	}

	return hex.EncodeToString(armored), nil
}

func EncryptSecretForProof(publicKey string, data []byte) (string, error) {
	// Parse the public key from PEM
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println("failed to decode PEM block containing public key")
		return "", errors.New("failed to decode PEM block containing public key")
	}

	// Parse the public key into rsa.PublicKey
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse public key: ", err)
		return "", errors.New("failed to parse public key")
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		fmt.Println("not an RSA public key")
		return "", errors.New("not an RSA public key")
	}

	// Hash for OAEP padding
	hash := sha256.New()

	// Encrypt the plaintext using RSA-OAEP with SHA-256
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPub, []byte(data), nil)
	if err != nil {
		fmt.Println("encryption failed: ", err)
		return "", errors.New("encryption failed")
	}

	// Convert to hex string (similar to OpenSSL's xxd -p)
	hexString := fmt.Sprintf("%x", ciphertext)

	fmt.Println("Encrypted Hex String:", hexString)
	return hexString, nil
}
