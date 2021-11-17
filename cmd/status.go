/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	b64 "encoding/base64"

	"github.com/karnowsa/gologic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of every or of a specific Server",
	Long: `The Weblogic REST API gives the opportunity to check the status of the diffrent type of server (AdminServer, ManagedServer, NodeManager)

With gologictl you can check the status, with the command 
"gologictl status" or if you want to check the status of 
a specific Managed Server you can pass a list of ManagedServers.
For example "gologictl status managed1 managedserver2..."`,
	Run: func(cmd *cobra.Command, args []string) {
		passwordBase64Decode, _ := b64.StdEncoding.DecodeString(viper.GetString("password"))
		var admin gologic.AdminServer = gologic.LoginAdminServer(
			viper.GetString("ip"),
			viper.GetInt("port"),
			viper.GetString("username"),
			string(passwordBase64Decode))
		admin.PrintStatus(args)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
