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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "The user login",
	Long: `Users should be logined with username, password.
	For example:
	agenda login -u=UserTest -p=10086
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Login Command:")
		loginFun(rName, rPassword)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVarP(&rName, "username", "u", "", "u=userName")
	loginCmd.Flags().StringVarP(&rPassword, "password", "p", "", "-p=password")
}

func loginFun(name string, password string) {
	logInit("Cmd login called")
	onUser := data.LoadOnUser()
	if onUser != nil {
		fmt.Println("You have to logout and then login a new user account!")
		logSave(name+"login failed while there is other on!", "[Error]")
		return
	}

	userSet := data.LoadUser()
	for _, userUnit := range userSet {
		if userUnit.GetUsername() == name {
			if userUnit.GetPassword() == password {
				fmt.Println(name + " Login successfully!")
				data.WriteOnUser(&userUnit)
				logSave(name+" login successfully", "[Info]")
				return
			}
			fmt.Println("login failed")
			logSave(name+"login failed while the password unmatches the user", "[Error]")
			return
		}
	}
	fmt.Println("Login with a unregistered user, failed!")
	logSave("Login failed with unregistered user", "[Error]")
}
