package tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/waymobetta/chainsafe-interview-2/ipfs"
)

// public function to test storing file in IPFS
func TestStore(t *testing.T) {
	// utilize go dot env package
	err := godotenv.Load("../.env")
	if err != nil {
		t.Log("[ipfs] error loading .env file")
		os.Exit(1)
	}

	// init new local temp file to write sample data to
	file, err := ioutil.TempFile(".", "")
	if err != nil {
		t.Errorf("[ipfs] %v\n", err)
	}

	// remove temp file before function termination
	defer os.Remove(
		file.Name(),
	)

	// write sample data to temp file
	if _, err := file.Write([]byte("foo")); err != nil {
		t.Errorf("[ipfs] %v\n", err)
	}

	// init new instance of Ipfs
	ipfsSvc := ipfs.New(
		&ipfs.Config{
			NodeUrl: os.Getenv("CS_IPFS_NODE_URL"),
		},
	)

	// set filename
	ipfsSvc.Filename = file.Name()

	// invoke storage of file
	if err := ipfsSvc.Store(); err != nil {
		t.Errorf("[ipfs] %v\n", err)
	}

	// check if hash was created and set
	if ipfsSvc.Cid == "" {
		t.Errorf("[ipfs] hash not found: %v\n", err)
	}

	t.Log(ipfsSvc.Cid)
}
