/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"

	"github.com/desperateofstruggle/agenda/data"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "The user logout",
	Long: `Users can be logout
	For example:
	agenda logout
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Logout Command:")
		logoutFun()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logoutFun() {
	logInit("Cmd logout called")
	onUser := data.LoadOnUser()
	if onUser == nil {
		logSave("No user login", "[Error]")
		fmt.Println("Now no user login")
		return
	}
	data.WriteOnUser(nil)
	fmt.Println(onUser.GetUsername() + " Logout Successfully!")
	logSave(onUser.GetUsername()+"logout successfully", "[Info]")
}
