package main

import (
	"encoding/json"
	"fmt"

	"github.com/alebaffa/wanikani-webapp/apiutils"
	"github.com/alebaffa/wanikani-webapp/structures"
)

func main() {

	var user structures.User
	jsonResponse, err := apiutils.FetchJson()

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonResponse, &user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user.UserInfo.Username, ",", user.UserInfo.Title)

}
