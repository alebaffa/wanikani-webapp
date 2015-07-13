package main

import (
	"fmt"

	"github.com/alebaffa/wanikani-webapp/apiutils"
)

func main() {

	user := apiutils.GetKanjiList("1")
	for _, character := range user.RequestedInfo {
		fmt.Printf("%#U", character.Character)
	}

}
