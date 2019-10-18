package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/desperateofstruggle/agenda/entity"
)

var UsersfilePath string = "./data/Users.json"
var OnUserfilePath string = "./data/curUser.txt"

// LoadUser ...
func LoadUser() []entity.User {
	var users []entity.User
	_, err := os.Stat(UsersfilePath)
	if err != nil {
		if !os.IsExist(err) {
			return users
		}
	}
	jsons, err := ioutil.ReadFile(UsersfilePath)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return users
	}
	strs := strings.Replace(string(jsons), "\n", "", 1)
	if strs == "" {
		return users
	}
	err = json.Unmarshal(jsons, &users)
	if err != nil {
		fmt.Println("Unmarshal of Json: ", err.Error())
		return users
	}
	return users
}

// WriteUser ...
func WriteUser(u []entity.User) {
	os.Truncate(UsersfilePath, 0)

	jsons, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Unmarshal of Json: ", err.Error())
	}
	err = ioutil.WriteFile(UsersfilePath, jsons, os.ModeAppend)
	if err != nil {
		fmt.Println("WriteFile: ", err.Error())
	}
	os.Chmod(UsersfilePath, 0777)

}

// LoadOnUser ...
func LoadOnUser() *entity.User {
	_, err := os.Stat(OnUserfilePath)
	if err != nil {
		if !os.IsExist(err) {
			return nil
		}
	}
	str, err := ioutil.ReadFile(OnUserfilePath)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil
	}
	strs := strings.Replace(string(str), "\n", "", 1)
	if strs == "" {
		return nil
	}
	userSet := LoadUser()
	for _, userUnit := range userSet {
		if userUnit.GetUsername() == strs {
			var tmps = new(entity.User)
			tmps.GainUser(userUnit)
			return tmps
		}
	}
	return nil
}

// WriteOnUser ...
func WriteOnUser(u *entity.User) {
	os.Truncate(OnUserfilePath, 0)
	if u == nil {
		return
	}

	err := ioutil.WriteFile(OnUserfilePath, []byte(u.GetUsername()), os.ModeAppend)
	if err != nil {
		fmt.Println("WriteFile: ", err.Error())
	}
	os.Chmod(OnUserfilePath, 0777)
}
