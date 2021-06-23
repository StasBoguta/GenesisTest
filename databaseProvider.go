package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)



func addNewUser(user User) bool{
	var users = readFile()

	if !isContains(users, user){
		users.Users = append(users.Users, user)
		file, _ := json.MarshalIndent(users, "", " ")
		_ = ioutil.WriteFile("credentials.json", file, 0644)
		return true
	}
	return false
}

func readFile() Users{
	jsonFile, _ := os.Open("credentials.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var users Users
	json.Unmarshal(byteValue, &users)
	return users
}

func isContains(users Users, user User) bool{
	for i:= 0; i < len(users.Users);i++{
		if users.Users[i].Login == user.Login{
			return true
		}
	}
	return false
}

func validateCredentials(user User) bool{
	users := readFile()
	for i:=0; i < len(users.Users); i++{
		if users.Users[i].Login == user.Login && users.Users[i].Password == user.Password{
			return true
		}
	}
	return false
}

