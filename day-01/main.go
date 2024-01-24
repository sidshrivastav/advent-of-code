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

func getFirstDigit(word string) string {
	wordLen := len(word)
	lowPtr := 0
	for lowPtr < wordLen {
		if unicode.IsDigit(rune(word[lowPtr])) {
			return string(word[lowPtr])
		}
		tempPtr := lowPtr
		numberWord := ""
		for tempPtr < wordLen {
			tempPtr += 1
			tempWord := word[lowPtr: tempPtr]
			numberWord = numberRepresentation[tempWord]
			if numberWord != "" {
				break
			}
		}

		if numberWord != "" {
			return numberWord
		}

		lowPtr += 1
	}

	return "0"
}

func reverseString(word string) (output string) {
	for _, character := range word {
		output = string(character) + output
	}

	return 
}

func getSecondDigit(word string) string {
	word = reverseString(word)
	wordLen := len(word)
	lowPtr := 0
	for lowPtr < wordLen {
		if unicode.IsDigit(rune(word[lowPtr])) {
			return string(word[lowPtr])
		}
		tempPtr := lowPtr
		numberWord := ""
		for tempPtr < wordLen {
			tempPtr += 1
			tempWord := word[lowPtr: tempPtr]
			tempWord = reverseString(tempWord)
			numberWord = numberRepresentation[tempWord]
			if numberWord != "" {
				break
			}
		}

		if numberWord != "" {
			return numberWord
		}

		lowPtr += 1
	}

	return "0"
}

func getNumberByParseWord(word string) int {
	numberBuilder := getFirstDigit(word) + getSecondDigit(word)
	number, err := strconv.Atoi(numberBuilder)
	if err != nil {
		panic(err)
	}
	return number

}

func solve(part string, wordList ...string) (result int) {
	for _, value := range wordList {
		if part == "1" {
			result += getNumber(value)
		}

		if part == "2" {
			result += getNumberByParseWord(value)
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
	
