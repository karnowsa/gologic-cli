/*
Copyright © 2020 karnowsa <karnowsa@mailbox.org>

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

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app [list of paths]",
	Args:  cobra.MinimumNArgs(1),
	Short: "Pass a target and a list of ",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetStringSlice("target")
		passwordBase64Decode, _ := b64.StdEncoding.DecodeString(viper.GetString("password"))
		var admin gologic.AdminServer = gologic.LoginAdminServer(
			viper.GetString("ip"),
			viper.GetInt("port"),
			viper.GetString("username"),
			string(passwordBase64Decode))
		admin.Deploy(target, args)
	},
}

func init() {
	var target []string

	deployCmd.AddCommand(appCmd)

	appCmd.Flags().StringSliceVarP(&target, "target", "t", []string{}, "Pass a list of targets (required)")
	appCmd.MarkFlagRequired("target")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
