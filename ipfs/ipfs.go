package ipfs

import (
	"bytes"
	"fmt"
	"io/ioutil"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

// Config is public strut to hold IPFS configurations
type Config struct {
	// URL of IPFS node
	NodeUrl string
}

// public struct to hold data for Ipfs
type Ipfs struct {
	// Filename is file to have contents be stored on IPFS
	Filename string
	// Cid is hash (string) receipt of stored data on IPFS
	Cid string
	// Config is configuration for IPFS node connection
	Config *Config
}

// public function to create a nw instance of Ipfs
// @param config is pointer to config struct
// @return pointer to newly created Ipfs struct
func New(config *Config) *Ipfs {
	// return pointer to newly created instance of Ipfs
	// Note: must be struct literal
	// new() function does not initiate nested pointers
	return &Ipfs{
		Config: config,
	}
}

// public method to handle storage of a file
// stores file data in IPFS
// sets CID
// @return error
func (i *Ipfs) Store() error {
	log.Printf("[ipfs] storing file: %s\n", i.Filename)

	// ipfs daemon
	sh := shell.NewShell(i.Config.NodeUrl)

	// read file contents here
	fileBytes, err := ioutil.ReadFile(i.Filename)
	if err != nil {
		return fmt.Errorf("[ipfs] %v", err)
	}

	// add file to IPFS
	cid, err := sh.Add(
		// read file bytes with new bytes reader
		bytes.NewReader(
			fileBytes,
		),
	)
	if err != nil {
		return fmt.Errorf("[ipfs] %v", err)
	}

	// set returned cid
	i.Cid = cid

	log.Printf("[ipfs] successfully stored file: %s -> %s\n", i.Filename, i.Cid)

	return nil
}
