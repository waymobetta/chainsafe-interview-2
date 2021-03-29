package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// Custom Errors
var errPublicKey = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")

// Config is public struct to handle config data for eth package
type Config struct {
	// URL of ethereum node (local; Ganache)
	ClientUrl string
	// private key -> environ var (security)
	PrivateKey string
	// storage contract address
	ContractAddress string
}

// Eth is public struct to hold data relating to storage of IPFS CID
type Eth struct {
	// Config is configuration for Eth transactions
	Config *Config
}

// public function to return new instance of Eth
// @param config is pointer to config struct
// @return pointer to newly initiated Eth struct
func New(config *Config) *Eth {
	// Note: must be struct literal
	// new() function does not initiate nested pointers
	return &Eth{
		Config: config,
	}
}

// public method to store CID of data stored in IPFS
// @param CID is string of IPFS storage receipt
// @return error
func (e *Eth) StoreCID(cid string) error {
	log.Printf("[eth] storing CID: %s\n", cid)

	// connect to ethereum node
	client, err := ethclient.Dial(e.Config.ClientUrl)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// convert private key to pointer to ECDSA private key
	privateKey, err := crypto.HexToECDSA(e.Config.PrivateKey)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// derive public key
	publicKey := privateKey.Public()

	// cast public key to pointer to ECDSA public key
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(errPublicKey)
	}

	// derive address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get next nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// query gas price to determine affordable rate
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// init chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// init new transactor from private key
	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		chainID,
	)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// convert uint64 to pointer to big int
	auth.Nonce = big.NewInt(
		int64(nonce),
	)

	// set value of tx -> $0
	auth.Value = big.NewInt(0) // in wei

	// set gas max
	auth.GasLimit = uint64(3000000) // in units

	// set gas price
	auth.GasPrice = gasPrice

	// convert contract string to common address type
	contractCommonAddress := common.HexToAddress(e.Config.ContractAddress)

	// init new instance of contract
	contractInstance, err := NewStore(
		contractCommonAddress,
		client,
	)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	tx, err := contractInstance.AddCID(
		auth,
		cid,
	)
	if err != nil {
		return fmt.Errorf("[eth] %v", err)
	}

	// log tx receipt
	log.Printf("[eth] transaction successful: %s\n", tx.Hash().Hex())

	// if transaction confirmation received
	if len(tx.Hash().Hex()) > 0 {
		log.Printf("[eth] successfully stored CID: %s\n", cid)
	}

	return nil
}

// public method to retrieve all CIDs stored in contract
// @return cids is slice of strings
// @return error
func (e *Eth) GetCIDs() (
	[]string,
	error,
) {
	log.Printf("[eth] retrieving CIDs from contract: %s\n",
		e.Config.ContractAddress,
	)

	// connect to ethereum node
	client, err := ethclient.Dial(e.Config.ClientUrl)
	if err != nil {
		return nil, fmt.Errorf("[eth] %v", err)
	}

	// convert contract string to common address type
	contractCommonAddress := common.HexToAddress(e.Config.ContractAddress)

	// init new instance of contract
	contractInstance, err := NewStore(
		contractCommonAddress,
		client,
	)
	if err != nil {
		return nil, fmt.Errorf("[eth] %v", err)
	}

	// get CID string array from contract
	// use custom function from contract rather than auto-generated
	// getter from state variable creation for easier use
	// takes required argument of call opts which we do not need -> nils
	cids, err := contractInstance.GetCIDArray(nil)
	if err != nil {
		return nil, fmt.Errorf("[eth] %v", err)
	}

	// return if no cids found
	if len(cids) == 0 {
		return nil, fmt.Errorf("[eth] no CIDs found in contract")
	}

	log.Printf("[eth] successfully returned [%v] CIDs\n", len(cids))

	return cids, nil
}
