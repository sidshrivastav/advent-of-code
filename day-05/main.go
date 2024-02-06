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

func getMapping(data [][]int) map[int]int {
	mapping := make(map[int]int)
	for _, value := range data {
		for idx := 0; idx < value[2]; idx++ {
			mapping[value[1]+idx] = value[0] + idx
		}
	}
	return mapping
}

func getMappingValue(data map[int]int, key int) (int) {
	val, ok := data[key]
	if ok {
		return val
	}

	return key
}

func getLowestLocation(data map[string][][]int) (int) {
	lowest := math.MaxInt32
	seedValues, ok := data["seeds"]
	if ok {
		delete(data, "seeds")
	}
	mapping := map[string]map[int]int{}
	for key, value := range data {
		mapping[key] = getMapping(value)
	}
	seeds := seedValues[0]
	for _, seed := range seeds {
		soil := getMappingValue(mapping["seed-to-soil map:"], seed)
		fertilizer := getMappingValue(mapping["soil-to-fertilizer map:"], soil)
		water := getMappingValue(mapping["fertilizer-to-water map:"], fertilizer)
		light := getMappingValue(mapping["water-to-light map:"], water)
		temp := getMappingValue(mapping["light-to-temperature map:"], light)
		humidity := getMappingValue(mapping["temperature-to-humidity map:"], temp)
		location := getMappingValue(mapping["humidity-to-location map:"], humidity)
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
	almanac := parseInput("input_1.txt")
	fmt.Println(getLowestLocation(almanac))
}
