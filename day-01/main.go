package main

import (
	"fmt"
	"unicode"
	"strconv"
	"os"
	"strings"
)

func getNumber(word string) (int) {
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

func solve(wordList ...string) (result int) {
	for _, value := range wordList {
		result += getNumber(value)
	}
	return 
}

func main() {
	inputFile := "input.txt"
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	wordList := strings.Split(string(content), "\n")
	result := solve(wordList...) 
	fmt.Println(result)
}
	
