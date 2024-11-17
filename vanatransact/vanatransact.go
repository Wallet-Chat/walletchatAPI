package vanatransact

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"rest-go-demo/vanaDataRegistryContract"
	"rest-go-demo/vanaDlpContract"
	"rest-go-demo/vanaTeeContract"

	"time"

	_ "rest-go-demo/docs"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetDlpPublicKey() string {
	// Connect to an vana mokshanode
	client, err := ethclient.Dial(os.Getenv("VANA_RPC_URL"))
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Create an instance of the contract
	instance, err := vanaDlpContract.NewVanaDlpContract(common.HexToAddress(os.Getenv("VANA_DLP_CONTRACT")), client)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	// Call the contract method
	var result string
	result, err = instance.MasterKey(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
		return "nil"
	}

	return result
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

	fmt.Println("Uploaded File ID: ", fileID)

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

func GetTeeDetails(latestJobId big.Int) {
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

	//return
	fmt.Println("Job File Result: ", getJobFileResult)
	fmt.Println("Tee Info Result: ", getTeeInfoResult)
}

func GetTeeContributionProof(fileId string) string {
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
	fileIdBigInt := new(big.Int)
	fileIdBigInt.SetString(fileId, 10)

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
	teePocTX, err := instance.RequestContributionProof(opts, fileIdBigInt)
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

	permissions := []vanaDataRegistryContract.IDataRegistryPermission{
		{Key: permissionsEEK, Account: ownerWallet},
	}

	// Send transaction
	tx, err := instance.AddFileWithPermissions(opts, encryptedFileUrl, ownerWallet, permissions)
	if err != nil {
		// Check revert reason
		callMsg := ethereum.CallMsg{
			From: opts.From,
			To:   &contractAddress,
			Data: tx.Data(),
		}
		result, callErr := client.CallContract(context.Background(), callMsg, nil)
		if callErr == nil && len(result) > 0 {
			return "", fmt.Errorf("transaction reverted: %s", string(result))
		}
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
