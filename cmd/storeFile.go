/*
Copyright © 2021 Jon Roethke <waymobetta@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/waymobetta/chainsafe-interview-2/eth"
	"github.com/waymobetta/chainsafe-interview-2/ipfs"
)

// storeFileCmd represents the storeFile command
var storeFileCmd = &cobra.Command{
	Use:     "store /path/to/file",
	Short:   "Store file on IPFS",
	Long:    `Store file on IPFS. Returns a CID that is then programmatically stored in a smart contract.`,
	Example: "usage: ./chainsafe-interview-2 store /path/to/file",

	Run: func(cmd *cobra.Command, args []string) {
		// panic if local file not passed as argument
		if len(os.Args) < 3 {
			// throw missing argument error
			if err := errors.New(
				"[cmd/storeFile] must pass file as argument\n",
			); err != nil {
				// show usage example
				fmt.Println(cmd.Example)
				log.Error(err)
				os.Exit(1)
			}
		}

		// full path of file to store
		filePath := os.Args[2]

		// init new instance of Ipfs
		// load .env -> github.com/waymobetta/chainsafe-interview-2/cmd/root.go
		// Note: environ vars should load natively from Viper (TODO)
		ipfsSvc := ipfs.New(
			&ipfs.Config{
				NodeUrl: os.Getenv("CS_IPFS_NODE_URL"),
			},
		)

		// set filename
		ipfsSvc.Filename = filePath

		// invoke storage of file
		if err := ipfsSvc.Store(); err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// init new instance of Eth
		// load .env -> github.com/waymobetta/chainsafe-interview-2/cmd/root.go
		// Note: environ vars should load natively from Viper (TODO)
		ethSvc := eth.New(
			&eth.Config{
				ClientUrl:       os.Getenv("CS_CLIENT_URL"),
				PrivateKey:      os.Getenv("CS_PRIVATE_KEY"),
				ContractAddress: os.Getenv("CS_CONTRACT_ADDRESS"),
			},
		)

		// store CID in contract
		err := ethSvc.StoreCID(ipfsSvc.Cid)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(storeFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
