package main 

import "fmt"
import "os"
import "strings"
import "math"


func getInput(inputFile string) ([]string) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(content), "\n")
	return input[:len(input) - 1]
}

func getPoints(cards []string) (points int){
	for _, card := range cards {
		allPoints := strings.Split(strings.Split(string(card), ":")[1], "|")
		givenNumbers := strings.Split(strings.TrimSpace(allPoints[0]), " ")
		winningNumbers := strings.Split(strings.TrimSpace(allPoints[1]), " ")
		numbersWon := 0
		for _, number := range givenNumbers {
			if number != "" {
				// TODO: Implement a binary search, just from fun
				for _, winningNumber := range winningNumbers {
					if number == winningNumber {
						numbersWon += 1
						break
					}
				}
			}
		}
		earned := 0
		if numbersWon > 0 {
			earned = int(math.Pow(2, float64(numbersWon - 1)))
		}
		points += earned
	}
	return
}

func getTotalScratchcards(card []string) (scratchcards int) {

	return 
}

func main() {
	// Part 1
	cards := getInput("input_1.txt")
	fmt.Println(int(getPoints(cards)))

	// Part 2
	cards = getInput("input_2.txt")
	fmt.Println(int)
}
