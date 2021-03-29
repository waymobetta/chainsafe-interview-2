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
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/waymobetta/chainsafe-interview-2/eth"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored CIDs",
	Long:  `List out the CIDs (IPFS hashes) that are currently stored in the contract.`,

	Run: func(cmd *cobra.Command, args []string) {
		// init new instance of Eth
		// load .env -> github.com/waymobetta/chainsafe-interview-2/cmd/root.go
		// Note: environ vars should load natively from Viper (TODO)
		ethSvc := eth.New(
			&eth.Config{
				ClientUrl:       os.Getenv("CS_CLIENT_URL"),
				ContractAddress: os.Getenv("CS_CONTRACT_ADDRESS"),
			},
		)

		// get all CIDs from contract
		cids, err := ethSvc.GetCIDs()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// loop over cid slice to pretty print
		for _, cid := range cids {
			fmt.Println(cid)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
