package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

var pointsMap = map[string]int{"Paper": 2, "Rock": 1, "Scissors": 3}
var countersMap = map[string]string{"Y": "A", "X": "C", "Z": "B"}

const winPoints = 6
const lossPoints = 0
const drawPoints = 3

func calculateRound(words []string) (int, error) {

	points, ok := pointsMap[words[1]]

	if !ok {
		return 0, errors.New("invalid option")
	}

	return points, nil

}

func calculateSum(filename string) (int, error) {

	sum := 0
	file, err := os.Open(filename)
	if err != nil {
		return sum, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		points, err := calculateRound(words)
		if err != nil {
			return sum, err
		}
		sum += points
	}

	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}

func main() {

	nPoints, err := calculateSum(`input.txt`)
	if err != nil {
		log.Panic(err)
	}
	log.Println(nPoints)
}
