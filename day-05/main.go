package main

import "fmt"
import "os"
import "strings"
import "strconv"

func stringToIntArray(arr []string) (result []int) {
	// Convert string values to integers
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return
}

func getInput(inputFile string) (map[string][][]int) {
	output := map[string][][]int{}
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(content), "\n\n")
	input = input[:len(input) - 1]
	
	// Parse Seeds
	seed := strings.Split(input[0], ":")
	seedValues := strings.Split(strings.TrimSpace(seed[1]), " ")

	output[seed[0]] = append(output[seed[0]], stringToIntArray(seedValues))

	// Handles Maps
	for idx, values := range input {
		if idx != 0 {
			valuesArray := strings.Split(string(values), "\n")
			for idx1, mappedValues := range valuesArray {
				if idx1 != 0 {
					temp := strings.Split(strings.TrimSpace(mappedValues), " ")
					output[valuesArray[0]] = append(output[valuesArray[0]], stringToIntArray(temp))
				}
			}
		}
	}
	return output
}

func main() {
	// Part 1
	almanac := getInput("input_1.txt")
	fmt.Println(almanac)
}
