package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Take in users input ensuring that more than one arguement is entered
	input := os.Args[1:]
	// returns nothing if one or less args entered
	if len(input) < 2 {
		return
	}
	// goes through each character in the input and checks if its outside bounds of
	// printable character(32-126), if so then print nothing
	// (this is because runes can range to unicode which has way more characters)
	for _, word := range input[0] {
		if word < 32 || word > 126 {
			return
		}
	}
	// This reads the standard.txt file and checks for error
	font := os.Args[2]
	bytes, err := ioutil.ReadFile(font + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// splits text file by new line
	//newAscii := strings.Split(string(bytes), "\n")
	var newAscii []string
	//have to treat thinkertoy different to standard and shadow
	if font == "thinkertoy" {
		//for some reason thinkertoy needs to be split by '\r\n' otherwise will print gibberish
		newAscii = strings.Split(string(bytes), "\r\n")

	} else {
		//cannot be split by '\r\n' otherwise will print gibberish
		newAscii = strings.Split(string(bytes), "\n")

	}
	//
	newInput := strings.Split(input[0], "\\n")
	var userInput []rune
	// creates blank array of runes to be appeneded with ascii equivalent of the users input
	for _, word := range newInput {
		if word == "" {
			fmt.Println()
			continue
		}
		for _, character := range word {
			if character == 92 && word[len(word)-2] == 92 {
				remove(newInput)
				continue
			}
			userInput = append(userInput, character)
		}
		printArtAscii(userInput, newAscii)
		userInput = []rune{}
	}
	// userInput = append(userInput[:i], userInput[i+1:] ...)
}

// users input and re-formatted standard.txt file passed as aruguments in pritnart function
func printArtAscii(userInput []rune, Ascii []string) {
	// loops through each row of individual ascii character in art file
	for line := 1; line <= 8; line++ {
		// Loop through each character of users input
		for _, character := range userInput {
			//Match users input(rune) with each row of Ascii Art
			skip := (character - 32) * 9
			// print the line from art file at the position specified
			//by calculation of skip to find the corresponding line in to users input
			fmt.Print(Ascii[line+int(skip)])
		}
		// Tells the function to skip to the next line before commencing next loop
		fmt.Println()
	}
}
func remove(input []string) []string {
	var empty []string
	for _, character := range input {
		if character == "\\" {
			newstr := fmt.Sprint(input)
			New := strings.Replace(newstr, "\\", "", -1)
			empty = append(empty, New)
		}
	}
	return empty
}
