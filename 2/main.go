package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
	"time"
)

type Result struct {
	Words []string
	Err error
}

func timeFunc(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type CalculatorFunc func ([]string) (int, error)

func calculateStrategy(words []string) (int, error) {

	opponent, ok := aliases[words[0]]

	if !ok {
		return 0, errors.New("invalid option")
	}

	outcome, ok := outcomes[words[1]]

	if !ok {
		return 0, errors.New("invalid outcome")
	}

	points := outcome

	switch outcome {
	case drawPoints:
		points += opponent.Points
	case lossPoints:
		losingGesture, ok := aliases[opponent.Beats[1]]
		if ! ok {
			return 0, errors.New("invalid gesture")
		}
		points += losingGesture.Points
	case winPoints:
		for _, g := range gestures{
			if g.Beats[0] == opponent.Name {
				points += g.Points
			}
		}

		if points == outcome {
			return points, errors.New("invalid gesture, can't be beaten")
		}

	}

	return points, nil


}

func calculateRound(words []string) (int, error) {

	opponent, ok := aliases[words[0]]

	if !ok {
		return 0, errors.New("invalid option")
	}

	player, ok := aliases[words[1]]

	if !ok {
		return 0, errors.New("invalid option")
	}

	points := player.Points

	if player.Beats[0] == opponent.Name {
		return points + winPoints, nil
	}

	if opponent.Beats[0] == player.Name {
		return points + lossPoints, nil
	}


	return points + drawPoints, nil

}

func calculateSum(filename string, calculator CalculatorFunc) (int, error) {

	defer timeFunc(time.Now(), "Calculator")
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

		points, err := calculator(words)
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

	sum, err := calculateSum(`input.txt`, calculateRound)
	if err != nil {
		log.Panic(err)
	}
	log.Println(sum)

	sum, err = calculateSum(`input.txt`, calculateStrategy)
	if err != nil {
		log.Panic(err)
	}
	log.Println(sum)

}
