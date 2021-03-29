package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/waymobetta/chainsafe-interview-2/eth"
)

// global IPFS hash for testing
var (
	testIPFSHASH = "QmdUWc45kg8QhuG5e86rVisaGHUMcVSG9DqkGhd8pq1HdK"
)

// StoreTest is public struct to hold config data for simulated chain
type StoreTest struct {
	auth    *bind.TransactOpts
	address common.Address
	gAlloc  core.GenesisAlloc
	sim     *backends.SimulatedBackend
	store   *eth.Store
}

// public function to handle setting up an instance of StoreTest
// handles testing all other tests internally
func TestSetup(t *testing.T) {
	// init new instance of StoreTest
	s := &StoreTest{
		auth:  &bind.TransactOpts{},
		sim:   &backends.SimulatedBackend{},
		store: &eth.Store{},
	}

	// generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("[eth] %v\n", err)
	}

	// init new auth transactor
	s.auth = bind.NewKeyedTransactor(privateKey)

	// set calling address (owner)
	s.address = s.auth.From

	// init balance
	bal := new(big.Int)

	// convert balance (string) to pointer to big Int
	// 10 ETH in Wei
	// https://eth-converter.com/
	bal.SetString("10000000000000000000", 10)

	// set account balance allocation at genesis
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {
			Balance: bal,
		},
	}

	// manually set gas limit
	gasLimit := uint64(3000000)

	// set and init simulated chain
	s.sim = backends.NewSimulatedBackend(s.gAlloc, gasLimit)

	// handle deploying contract
	s.testDeployContract(t)

	// handle storing CID in contract
	s.testStoreCID(t)

	// handle retrieving stored CIDs from contract
	s.testGetCIDs(t)
}

// private function to test deploying a contract to a simulated chain
func (s *StoreTest) testDeployContract(t *testing.T) {
	// deploy contract
	contractAddress, _, store, err := eth.DeployStore(s.auth, s.sim)
	if err != nil {
		t.Errorf("[eth] %v\n", err)
	}

	// set store instance
	s.store = store

	// commit transactions to state
	s.sim.Commit()

	t.Logf("[eth] contract successfully deployed at: %s\n",
		contractAddress.String(),
	)
}

// private function to test storing CID in contract on simulated chain
func (s *StoreTest) testStoreCID(t *testing.T) {
	// add CID in contract
	tx, err := s.store.AddCID(
		s.auth,
		testIPFSHASH,
	)
	if err != nil {
		t.Error(err)
	}

	t.Logf("[eth] transaction successful: %s\n", tx.Hash().Hex())

	// commit new block manually
	s.sim.Commit()

	// if transaction confirmation received
	if len(tx.Hash().Hex()) > 0 {
		t.Logf("[eth] successfully stored CID: %s\n", testIPFSHASH)
	}

	// retrieve tx receipt
	receipt, err := s.sim.TransactionReceipt(
		context.Background(),
		tx.Hash(),
	)
	if err != nil {
		t.Errorf("[eth] %v\n", err)
	}

	// if receipt type returns empty
	if receipt == nil {
		t.Errorf("[eth] receipt not found: %v\n", err)
	}

	// 1 if success
	t.Logf("[eth] transaction status: %d\n", receipt.Status)
}

// private function to test retrieving CIDs in contract on simulated chain
func (s *StoreTest) testGetCIDs(t *testing.T) {
	// retrieve CIDs from contract
	cids, err := s.store.GetCIDArray(nil)
	if err != nil {
		t.Errorf("[eth] %v\n", err)
	}

	// return if no cids found
	if len(cids) == 0 {
		t.Logf("[eth] no CIDs found in contract\n")
		return
	}

	t.Logf("[eth] successfully returned [%v] CIDs\n", len(cids))
}
