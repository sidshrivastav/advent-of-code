package main

import (
	"fmt"
	"unicode"
	"os"
	"strconv"
	"strings"
)

func getSum(row int, engineRow string, engine []string) int {
	start := 0
	rowSum := 0
	position := [][]int{}
	itr := 0
	engineRowLen := len(engineRow)
	end := engineRowLen - 1
	engineLen := len(engine)
	for itr < engineRowLen {
		if !unicode.IsDigit(rune(engineRow[itr])) {
			itr += 1
			continue
		}
		partPosition := []int{}
		partPosition = append(partPosition, itr)
		for itr < engineRowLen && unicode.IsDigit(rune(engineRow[itr])) {
			itr += 1
		}
		partPosition = append(partPosition, itr-1)
		position = append(position, partPosition)
	}
	
	for _, coords := range position {
		locationToCheck := [][]int{}
		numberStr := ""
		for itr := coords[0]; itr <= coords[1]; itr++ {
			numberStr += string(engineRow[itr])
			locationToCheck = append(locationToCheck, []int{row-1, itr})
			locationToCheck = append(locationToCheck, []int{row+1, itr})
			if itr != start && itr == coords[0] {
				locationToCheck = append(locationToCheck, []int{row-1, itr-1})
				locationToCheck = append(locationToCheck, []int{row+1, itr-1})
				locationToCheck = append(locationToCheck, []int{row, itr-1})
			}
			if itr != end && itr == coords[1] {
				locationToCheck = append(locationToCheck, []int{row-1, itr+1})
				locationToCheck = append(locationToCheck, []int{row+1, itr+1})
				locationToCheck = append(locationToCheck, []int{row, itr+1})
			}
		}
		validPart := false
		for _, location := range locationToCheck {
			if location[0] == 0 {
				continue
			}
			if location[0] == engineLen - 1 {
				continue
			}
			part :=  string(engine[location[0]][location[1]])
			// if !unicode.IsDigit(part) && string(part) != "." {
			if part != "." {
				validPart = true
			} 
		}

		if validPart {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			rowSum += number
		}
	}
	return rowSum
}

func getRatioSum(engine []string) (result int) {
	itr := 1
	end := len(engine) - 1
	starPosition := [][]int{}
	partPosition := [][]int{}
	for itr < end {
		itr2 := 0
		engineRowLen := len(engine[itr])
		for itr2 < engineRowLen {
			if string(engine[itr][itr2]) == "*" {
				position := []int{itr, itr2}
				starPosition = append(starPosition, position)
			}
			if !unicode.IsDigit(rune(engine[itr][itr2])) {
				itr2 += 1
				continue
			}
			position := []int{}
			position = append(position, itr)
			position = append(position, itr2)
			for itr2 < engineRowLen && unicode.IsDigit(rune(engine[itr][itr2])) {
				itr2 += 1
			}
			position = append(position, itr2-1)
			partPosition = append(partPosition, position)
		}
		itr += 1
	}
	ratio := map[int][]int{}
	for _, coords := range partPosition {
		locationToCheck := [][]int{}
		numberStr := ""
		row := coords[0]
		for itr := coords[1]; itr <= coords[2]; itr++ {
			numberStr += string(engine[row][itr])
			locationToCheck = append(locationToCheck, []int{row-1, itr})
			locationToCheck = append(locationToCheck, []int{row+1, itr})
			if itr != 0 && itr == coords[1] {
				locationToCheck = append(locationToCheck, []int{row-1, itr-1})
				locationToCheck = append(locationToCheck, []int{row+1, itr-1})
				locationToCheck = append(locationToCheck, []int{row, itr-1})
			}
			if itr != end && itr == coords[2] {
				locationToCheck = append(locationToCheck, []int{row-1, itr+1})
				locationToCheck = append(locationToCheck, []int{row+1, itr+1})
				locationToCheck = append(locationToCheck, []int{row, itr+1})
			}
		}
		
		for _, loc := range locationToCheck {
			for idx, sloc := range starPosition {
				if loc[0] == sloc[0] && loc[1] == sloc[1] {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						panic(err)
					}
					ratio[idx] = append(ratio[idx], number)
				}
			}
		}
	}
	
	for _, value := range ratio {
		if len(value) == 2 {
			result += value[0] * value[1]
		}
			
	}

	return
}


func solve(part string, engine []string) (result int) {
	engineLen := len(engine)
	if part == "1" {
		itr := 1
		end := engineLen - 1
		for itr < end {
			result += getSum(itr, engine[itr], engine)
			itr += 1
		}
	}
	if part == "2" {
		result = getRatioSum(engine)
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
	engine := strings.Split(string(content), "\n")
	engine = engine[:len(engine) - 1]
	column := len(engine[0])
	bufferRow := ""
	for itr := 0; itr < column; itr++ {
		bufferRow += "."
	}
	engine = append([]string{bufferRow}, engine...)
	engine = append(engine, bufferRow)
	result := solve("1", engine)
	fmt.Println(result)

	// Part 1
	inputFile = "input_2.txt"
	content, err = os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	engine = strings.Split(string(content), "\n")
	engine = engine[:len(engine) - 1]
	column = len(engine[0])
	bufferRow = ""
	for itr := 0; itr < column; itr++ {
		bufferRow += "."
	}
	engine = append([]string{bufferRow}, engine...)
	engine = append(engine, bufferRow)
	result = solve("2", engine)
	fmt.Println(result)

}	
