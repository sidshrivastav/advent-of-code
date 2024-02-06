package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "math"

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

func searchIndex(source int, length int, key int) int {
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if source + mid > key {
			high -= 1
		} else {
			low += 1
		}
	}
	return high
}

func getMappingValue(data [][]int, key int) (int) {
	for _, val := range data {
		if val[1] <= key && key < val[1] + (val[2]) {
			idx := searchIndex(val[1], val[2], key)
			if idx != -1 {
				return val[0] + idx
			}
		} 
	}
	return key
}

func getLowestLocation(data map[string][][]int) (int) {
	lowest := math.MaxInt32
	for _, seed := range data["seeds"][0] {
		soil := getMappingValue(data["seed-to-soil map:"], seed)
		fertilizer := getMappingValue(data["soil-to-fertilizer map:"], soil)
		water := getMappingValue(data["fertilizer-to-water map:"], fertilizer)
		light := getMappingValue(data["water-to-light map:"], water)
		temp := getMappingValue(data["light-to-temperature map:"], light)
		humidity := getMappingValue(data["temperature-to-humidity map:"], temp)
		location := getMappingValue(data["humidity-to-location map:"], humidity)
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func parseInput(inputFile string) (map[string][][]int) {
	output := map[string][][]int{}
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(content), "\n\n")
	input[len(input)-1] = input[len(input)-1][:len(input[len(input)-1]) - 1]
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
	almanac := parseInput("input_1.txt")
	fmt.Println(getLowestLocation(almanac))
}
