package main

import (
	"fmt"

	"github.com/alebaffa/wanikani-webapp/apiutils"
)

func main() {

	kanjin1 := apiutils.GetKanjiList("1")
	kanjin2 := apiutils.GetKanjiList("2")
	kanjin3 := apiutils.GetKanjiList("3")
	kanjin4 := apiutils.GetKanjiList("4")

	fmt.Println("\nKanji n1: ")
	for _, character := range kanjin1.RequestedInfo {
		fmt.Printf(character.Character + ", ")
	}

	fmt.Println("\nKanji n2: ")
	for _, character := range kanjin2.RequestedInfo {
		fmt.Printf(character.Character + ", ")
	}

	fmt.Println("\nKanji n3: ")
	for _, character := range kanjin3.RequestedInfo {
		fmt.Printf(character.Character + ", ")
	}

	fmt.Println("\nKanji n4: ")
	for _, character := range kanjin4.RequestedInfo {
		fmt.Printf(character.Character + ", ")
	}

}
