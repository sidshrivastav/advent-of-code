package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

func isConfigurationPossible(game string) int {
	red := 12 
	green := 13
	blue := 14
	gameId := strings.Split(strings.Split(game, ":")[0], " ")[1]
	gameData := strings.Split(strings.Split(game, ":")[1], ";")
	for _, gameSet :=  range gameData {
		for _, gameCubes := range strings.Split(gameSet, ",") {
			gameCube := strings.Split(gameCubes, " ")
			cubeCount, err := strconv.Atoi(gameCube[1])
			if err != nil {
				panic(err)
			}
			if gameCube[2] == "red" {
				if cubeCount > red {
					return 0
				}
			}
			if gameCube[2] == "green" {
				if cubeCount > green {
					return 0
				}
			}
			if gameCube[2] == "blue" {
				if cubeCount > blue {
					return 0
				}
			}
		}
	}
	

	gameIdNumber, err := strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	return gameIdNumber

}

func intMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func cubeSetPower(game string) int {
	red := 0 
	green := 0
	blue := 0
	gameData := strings.Split(strings.Split(game, ":")[1], ";")
	for _, gameSet :=  range gameData {
		for _, gameCubes := range strings.Split(gameSet, ",") {
			gameCube := strings.Split(gameCubes, " ")
			cubeCount, err := strconv.Atoi(gameCube[1])
			if err != nil {
				panic(err)
			}
			if gameCube[2] == "red" {
				red = intMax(cubeCount, red)
			}
			if gameCube[2] == "green" {
				green = intMax(cubeCount, green)
			}
			if gameCube[2] == "blue" {
				blue = intMax(cubeCount, blue)
			}
		}
	}
	

	return red * blue * green
}


func solve(part string, gameRound []string) (result int) {
	for _, value := range gameRound {
		if part == "1" && value != "" {
			result += isConfigurationPossible(value)
		}
		if part == "2" && value != "" {
			result += cubeSetPower(value)
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
	gameRound := strings.Split(string(content), "\n")
	fmt.Println("Part 1: ", solve("1", gameRound))

	// Part 1
	inputFile = "input_2.txt"
	content, err = os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	gameRound = strings.Split(string(content), "\n")
	fmt.Println("Part 2: ", solve("2", gameRound))

}
