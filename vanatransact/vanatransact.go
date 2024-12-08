package vanatransact

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"rest-go-demo/vanaDataRegistryContract"
	"rest-go-demo/vanaDlpContract"
	"rest-go-demo/vanaTeeContract"
	"rest-go-demo/vanaencrypt"
	"strings"

	"bytes"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	_ "rest-go-demo/docs"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetDlpPublicKey() (string, error) {
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Create an instance of the contract
	instance, err := vanaDlpContract.NewVanaDlpContract(common.HexToAddress(os.Getenv("VANA_DLP_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Call the contract method
	var result string
	result, err = instance.PublicKey(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result, nil
}

func GetFileID(txHash string) string {
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Wait for the transaction to be confirmed
	var fileID string
	for {
		// Check the transaction receipt
		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			// If the transaction is not yet mined, continue polling
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		// Check if the receipt is nil or has no logs
		if receipt == nil || len(receipt.Logs) == 0 {
			fmt.Println("Transaction failed or has no logs.")
			return "nil" // Return nil if the transaction failed
		}

		// Transaction is confirmed, extract the fileID from the receipt
		fileID = receipt.Logs[0].Topics[1].Hex() // Adjust this based on your contract's event
		break
	}

	// Convert hex string to big.Int
	fileIDBigInt := new(big.Int)
	fileIDBigInt.SetString(fileID[2:], 16) // Skip the "0x" prefix
	// Convert big.Int to string representation of the integer
	fileIDint := fileIDBigInt.String() // This will give you the decimal string representation
	fmt.Println("Uploaded File ID: ", fileIDint)

	return fileID // Return the extracted fileID
}

func GetTeePrice() string {
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Create an instance of the contract
	instance, err := vanaTeeContract.NewVanaTeeContract(common.HexToAddress(os.Getenv("VANA_TEE_POOL_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Create a call options object without a signer
	callOpts := &bind.CallOpts{
		Pending: false,            // Set to true if you want to query the pending state
		From:    common.Address{}, // Optional: specify the address if needed
	}

	// Call the contract method
	var result string
	teeFee, err := instance.TeeFee(callOpts)
	if err != nil {
		fmt.Println("tee getPrice fee error: ", err)
		return "nil"
	}
	result = teeFee.String()

	return result
}

func GetFileJobIDs(fileId string) []*big.Int {
	//fmt.Println("get file job ids for fileId: ", fileId)
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return []*big.Int{}
	}

	// Create an instance of the contract
	instance, err := vanaTeeContract.NewVanaTeeContract(common.HexToAddress(os.Getenv("VANA_TEE_POOL_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
		return []*big.Int{}
	}

	// Create a call options object without a signer
	callOpts := &bind.CallOpts{
		Pending: false,            // Set to true if you want to query the pending state
		From:    common.Address{}, // Optional: specify the address if needed
	}

	// Assuming fileId is a string representation of a number
	fileIdBigInt := new(big.Int)
	fileIdBigInt.SetString(fileId, 10)

	// Call the contract method
	getJobFileResult, err := instance.FileJobIds(callOpts, fileIdBigInt)
	if err != nil {
		fmt.Println("tee getFileJobIDs fee error: ", err)
		return []*big.Int{}
	}

	return getJobFileResult
}

func GetTeeDetails(latestJobId big.Int) (string, string) {
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
	}

	// Create an instance of the contract
	instance, err := vanaTeeContract.NewVanaTeeContract(common.HexToAddress(os.Getenv("VANA_TEE_POOL_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
	}

	// Create a call options object without a signer
	callOpts := &bind.CallOpts{
		Pending: false,            // Set to true if you want to query the pending state
		From:    common.Address{}, // Optional: specify the address if needed
	}

	// Call the contract method
	getJobFileResult, err := instance.Jobs(callOpts, &latestJobId)
	if err != nil {
		fmt.Println("tee getJobFileResult error: ", err, latestJobId)
	}

	// Call the contract method
	getTeeInfoResult, err := instance.Tees(callOpts, getJobFileResult.TeeAddress)
	if err != nil {
		fmt.Println("tee getTeeInfoResult error: ", err)
	}

	fmt.Println("Job File Result: ", getJobFileResult)
	fmt.Println("Tee Info Result: ", getTeeInfoResult)
	//return
	return getTeeInfoResult.Url, getTeeInfoResult.PublicKey
}

func GetTeeContributionProof(fileId string) string {
	//fmt.Println("get TEE contribution proof for fileId: ", fileId)
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Create an instance of the contract
	instance, err := vanaTeeContract.NewVanaTeeContract(common.HexToAddress(os.Getenv("VANA_TEE_POOL_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Assuming fileId is a string representation of a number
	// Convert hex string to big.Int
	fileIDBigInt := new(big.Int)
	fileIDBigInt.SetString(fileId[2:], 16) // Skip the "0x" prefix

	privateKey, err := crypto.HexToECDSA(os.Getenv("VANA_SIGNER_PRIVATE_KEY"))
	if err != nil {
		fmt.Println("invalid private key: %w", err)
		return ""
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(14800))
	if err != nil {
		fmt.Println("failed to create transactor: %w", err)
		return ""
	}

	// Call the contract method
	var result string
	fmt.Println("RequestContributionProof Proof for File: ", fileId, fileIDBigInt)
	teePocTX, err := instance.RequestContributionProof(opts, fileIDBigInt)
	if err != nil {
		fmt.Println("tee getContributionProof error: ", err)
		return "nil"
	}
	result = teePocTX.Hash().Hex()

	return result
}

func AddFileWithPermissions(ownerWallet common.Address, encryptedFileUrl string, permissionsEEK string) (string, error) {
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to connect to RPC: %w", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("VANA_DATA_REGISTRY_CONTRACT"))
	instance, err := vanaDataRegistryContract.NewVanaDataRegistryContract(contractAddress, client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("VANA_SIGNER_PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(14800))
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	//Account should be the DLP contract address since we're giving permission to the DLP to decrypt
	permissions := []vanaDataRegistryContract.IDataRegistryPermission{
		{Key: permissionsEEK, Account: common.HexToAddress(os.Getenv("VANA_DLP_CONTRACT"))},
	}

	// Send transaction
	tx, err := instance.AddFileWithPermissions(opts, encryptedFileUrl, ownerWallet, permissions)
	if err != nil {
		// Check revert reason
		// callMsg := ethereum.CallMsg{
		// 	From: opts.From,
		// 	To:   &contractAddress,
		// 	Data: tx.Data(),
		// }
		// result, callErr := client.CallContract(context.Background(), callMsg, nil)
		// if callErr == nil && len(result) > 0 {
		// 	return "", fmt.Errorf("transaction reverted: %s", string(result))
		// }
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	// Fetch receipt to ensure success
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return tx.Hash().Hex(), fmt.Errorf("failed to fetch receipt: %w", err)
	}

	if receipt.Status == 0 {
		return tx.Hash().Hex(), errors.New("transaction failed with status 0")
	}

	return tx.Hash().Hex(), nil
}

func convertHexStringToBigInt(fileId string) (*big.Int, error) {
	// Trim whitespace
	fileId = strings.TrimSpace(fileId)

	// Remove "0x" prefix if present
	if strings.HasPrefix(fileId, "0x") {
		fileId = fileId[2:]
	}

	// Convert hexadecimal string to big.Int
	fileIdInt, ok := new(big.Int).SetString(fileId, 16)
	if !ok {
		return nil, fmt.Errorf("failed to convert fileId to big.Int: %s", fileId)
	}

	return fileIdInt, nil
}

func RequestRewardFromDLP(fileId string) (string, error) {
	bigFileID, err := convertHexStringToBigInt(fileId)
	if err != nil {
		return "", fmt.Errorf("failed to convert fileID to bigint %w", err)
	}
	fmt.Println("Requesting Reward for FileID: ", bigFileID)
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to connect to RPC: %w", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("VANA_DLP_CONTRACT"))
	instance, err := vanaDlpContract.NewVanaDlpContract(contractAddress, client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("VANA_SIGNER_PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(14800))
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Send transaction
	tx, err := instance.RequestReward(opts, bigFileID, big.NewInt(1))
	if err != nil {
		fmt.Println("request reward err", err, opts.From, tx)
		// // Check revert reason
		// callMsg := ethereum.CallMsg{
		// 	From: opts.From,
		// 	To:   &contractAddress,
		// 	Data: tx.Data(),
		// }
		// result, callErr := client.CallContract(context.Background(), callMsg, nil)
		// if callErr == nil && len(result) > 0 {
		// 	return "", fmt.Errorf("transaction reverted: %s", string(result))
		// }
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	// Fetch receipt to ensure success
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return tx.Hash().Hex(), fmt.Errorf("failed to fetch receipt: %w", err)
	}

	if receipt.Status == 0 {
		return tx.Hash().Hex(), errors.New("transaction failed with status 0")
	}
	fmt.Println("Received Reward from DLP tx hash: ", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}

// Define the structure for the request body
type RequestBody struct {
	JobID                  *big.Int                 `json:"job_id"`
	FileID                 *big.Int                 `json:"file_id"`
	Nonce                  string                   `json:"nonce"`
	ProofURL               string                   `json:"proof_url"`
	EncryptionSeed         string                   `json:"encryption_seed"`
	EnvVars                map[string]string        `json:"env_vars"`
	Secrets                map[string]string        `json:"secrets"`
	ValidatePermissions    []ValidatePermissionTest `json:"validate_permissions"`
	EncryptedEncryptionKey string                   `json:"encrypted_encryption_key,omitempty"`
	EncryptionKey          string                   `json:"encryption_key,omitempty"`
}

// Define the structure for validate permissions
// type ValidatePermission struct {
// 	Address      string `json:"address"`
// 	PublicKey    string `json:"public_key"`
// 	IV           string `json:"iv"`
// 	EphemeralKey string `json:"ephemeral_key"`
// }

// Define the structure for validate permissions
type ValidatePermissionTest struct {
	Address      string `json:"address"`
	PublicKey    string `json:"public_key"`
	IV           string `json:"iv"`
	EphemeralKey string `json:"ephemeral_key"`
}

// Function to send the contribution proof request
func SendContributionProof(jobID *big.Int, fileID string, dlpPubKey string, envVars map[string]string, secrets map[string]string, teePublicKey string, teeURL string, iv []byte, emphemKeyPrivBytes []byte, signature string) error {
	// Create the request body

	// Convert hex string to big.Int
	fileIDBigInt := new(big.Int)
	fileIDBigInt.SetString(fileID[2:], 16) // Skip the "0x" prefix
	// Convert big.Int to string representation of the integer
	//fileIDint := fileIDBigInt.String()

	// Initialize the request body with empty ValidatePermissions
	requestBody := RequestBody{
		JobID:               jobID,
		FileID:              fileIDBigInt,
		Nonce:               "1234",
		ProofURL:            "https://github.com/vana-com/vana-satya-proof-template/releases/download/v24/gsc-my-proof-24.tar.gz",
		EncryptionSeed:      os.Getenv("VANA_ENCRYPT_KEY_SEED"),
		EnvVars:             envVars,
		Secrets:             secrets,
		ValidatePermissions: nil, // Initialize as an empty slice
	}

	// If TEE public key is available, encrypt the encryption key
	if teePublicKey != "" {
		fmt.Println("Encrypting encryption key with TEE public key: ", teePublicKey)
		//test only:
		//teePublicKey = "0xad1013116ea75ceb61b9f13a55cff6937a807b8e575dc2e2ccf7c1c115eab9d046e4a6507e235f3496481712c9a5be1fd37fcd018a70140b255a1abb16d9c678"

		//TODO, encryption phrase here is actually user signature which we should pass in as well not hard code
		encryptedKey, ephemeralKeyPriv, err := vanaencrypt.EncryptWithWalletPublicKey(signature, teePublicKey, iv, emphemKeyPrivBytes) // Implement this function

		// Set ValidatePermissions after encryptKey is populated
		requestBody.ValidatePermissions = []ValidatePermissionTest{
			{
				Address:      os.Getenv("VANA_DLP_CONTRACT"),
				PublicKey:    dlpPubKey,
				IV:           hex.EncodeToString(encryptedKey["iv"]),
				EphemeralKey: hex.EncodeToString(ephemeralKeyPriv),
			},
		}

		if err != nil {
			fmt.Println("Error encrypting encryption key:", err)
			fmt.Println("Warning: Failed to encrypt encryption key, falling back to direct encryption key")
			requestBody.EncryptionKey = signature
		} else {
			finalDataTeeEEK := append(encryptedKey["iv"],
				append(encryptedKey["ephemPublicKey"],
					append(encryptedKey["ciphertext"], encryptedKey["mac"]...)...)...)

			// Return the final result as a hex string
			hexDataTeeEEK := hex.EncodeToString(finalDataTeeEEK)
			requestBody.EncryptedEncryptionKey = hexDataTeeEEK
			fmt.Println("Encryption key encrypted successfully for TEE: ", hexDataTeeEEK)
		}
	} else {
		fmt.Println("TEE public key not available, using direct encryption key")
		requestBody.EncryptionKey = signature
	}
	//requestBody.EncryptionKey = signature //TODO - need to debug above code with Vana team

	fmt.Println("Sending contribution proof request to TEE")

	//test only:
	//testReqBody, err := CreateRequestBody()

	body, err := json.MarshalIndent(requestBody, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}
	fmt.Println("sending to RunProof: ", string(body))

	// Make the HTTP POST request
	client := &http.Client{Timeout: 600 * time.Second}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RunProof", teeURL), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Convert the body to a string and print it
	bodyString := string(bodyBytes)
	fmt.Println("**** /RunProof Response Body:", bodyString)

	if resp.StatusCode != http.StatusOK {
		var errorData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorData); err != nil {
			return fmt.Errorf("TEE request failed with status %d", resp.StatusCode)
		}
		return fmt.Errorf("TEE request failed: %v", json.NewDecoder(resp.Body))
	}

	return nil
}
