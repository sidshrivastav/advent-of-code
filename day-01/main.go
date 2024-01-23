package main

import (
	"fmt"
	"unicode"
	"strconv"
	"os"
	"strings"
)

var numberRepresentation = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

func getNumber(word string) int {
	wordLen := len(word)
	lowPtr := 0
	highPtr := wordLen - 1
	numberBuilder := ""
	for lowPtr < wordLen && !unicode.IsDigit(rune(word[lowPtr])) {
		lowPtr += 1
	}
	
	if lowPtr < wordLen {
		numberBuilder += string(word[lowPtr])
	}

	for highPtr >= 0 && !unicode.IsDigit(rune(word[highPtr])) {
		highPtr -= 1
	}

	if highPtr >= 0 {
		numberBuilder += string(word[highPtr])
	}
	if numberBuilder == "" {
		return 0
	}
	
	number, err := strconv.Atoi(numberBuilder)
	if err != nil {
		panic(err)
	}
	return number
}


func parseWord(word string) (normalWordBuilder string) {
	wordLen := len(word)
	lowPtr := 0
	for lowPtr < wordLen {
		tempPtr := lowPtr
		foundNumber := false
		for tempPtr < wordLen {
			tempPtr += 1
			tempWord := word[lowPtr: tempPtr]
			numberWord := numberRepresentation[tempWord]
			if numberWord != "" {
				normalWordBuilder += numberWord
				foundNumber = true
				break
			}
		}

		if foundNumber == true {
			lowPtr = tempPtr
			continue
		}

		normalWordBuilder += string(word[lowPtr])
		lowPtr += 1
	}

	//highPtr := wordLen - 1
	//for highPtr >= 0 {
	//	tempPtr := highPtr
	//	foundNumber := false
	//	for tempPtr >= 0 {
	//		tempPtr -= 1
	//		tempWord := word[tempPtr: highPtr+1]
	//
	//	}
	//}

	return 
}

func solve(part string, wordList ...string) (result int) {
	for _, value := range wordList {
		if part == "1" {
			result += getNumber(value)
		}

		if part == "2" {
			result += getNumber(parseWord(value))
		}
	}
	return 
}

func main() {
	// Part 1
	inputFile := "input_1.txt"
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	wordList := strings.Split(string(content), "\n")
	result := solve("1", wordList...) 
	fmt.Println(result)

	// Part 2
	inputFile = "input_2.txt"
	content, err = os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	wordList = strings.Split(string(content), "\n")
	result = solve("2", wordList...)
	fmt.Println(result)
}
	
