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
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/desperateofstruggle/agenda/data"
	"github.com/desperateofstruggle/agenda/entity"
)

var rName, rPassword, rEmai, rPhone string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new User",
	Long: `Users should be registered with username, password, phone and email.
	For example:
	agenda register -u=UserTest -p=10086 -e=kk@123.com -phone=13312345678
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		register(rName, rPassword, rEmai, rPhone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&rName, "username", "u", "", "-u=userName")
	registerCmd.Flags().StringVarP(&rPassword, "password", "p", "", "-p=password")
	registerCmd.Flags().StringVarP(&rEmai, "email", "e", "", "-e=email")
	registerCmd.Flags().StringVarP(&rPhone, "phone", "P", "", "-P=phone")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var infoLog *log.Logger
var logFile *os.File

func logInit(str string) {
	fileName := "./data/Agenda.log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

	if err != nil {
		log.Fatalln("Open file error")
	}
	infoLog = log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println(str)
}
func logSave(str string, logType string) {
	//fmt.Println(str)
	infoLog.SetPrefix(logType)
	infoLog.Println(str)
}

func register(name string, password string, email string, phone string) {
	logInit("Cmd register called")
	defer logFile.Close()
	if !ValidationOfName(name) || !ValidationOfPassword(password) || !ValidationOfEmail(email) ||
		!ValidationOfPhone(phone) {
		fmt.Println("format : agenda register -u [username] -p [password] -e [email] -P [phone]")
	} else {
		users := data.LoadUser()
		if entity.IsUserExisted(name, users) {
			logSave("The username has been registered", "[Warning]")
			logSave("Register fail", "[Warning]")
			fmt.Println("The username has been registered")
			return
		}
		var tmps = new(entity.User)
		tmps.Init(name, password, email, phone)
		users = append(users, *tmps)
		data.WriteUser(users)
		logSave("Register Successfully", "[Info]")
		fmt.Println("Register Successfully")
		return
	}
	logSave("Register fail", "[Warning]")
}

// ValidationOfName ...
func ValidationOfName(name string) bool {
	tmp := []byte(name)
	//fmt.Println(name)
	temp, _ := regexp.Match(".+", tmp)
	if !temp {
		logSave("flag -u ,name is invaild in format", "[Warning]")
	}
	return temp
}

// ValidationOfPassword ...
func ValidationOfPassword(pass string) bool {
	tmp := []byte(pass)
	temp, _ := regexp.Match(".+", tmp)
	if !temp {
		logSave("flag -p ,password is invaild in format", "[Warning]")
	}
	return temp
}

// ValidationOfEmail ...
func ValidationOfEmail(email string) bool {
	tmp := []byte(email)
	temp, _ := regexp.Match("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", tmp)
	if !temp {
		logSave("flag -e ,email is invaild in format", "[Warning]")
	}
	return temp
}

// ValidationOfPhone ...
func ValidationOfPhone(phone string) bool {
	tmp := []byte(phone)
	temp, _ := regexp.Match("^(13[0-9]|14[5-9]|15[0-3,5-9]|16[2,5,6,7]|17[0-8]|18[0-9]|19[1,3,5,8,9])\\d{8}$", tmp)
	if !temp {
		logSave("flag -P ,phone is invaild in format", "[Warning]")
	}
	return temp
}
