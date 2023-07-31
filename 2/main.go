package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Result struct {
	Words []string
	Err   error
}

func timeFunc(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type CalculatorFunc func([]string) (int, error)

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
		if !ok {
			return 0, errors.New("invalid gesture")
		}
		points += losingGesture.Points
	case winPoints:
		for _, g := range gestures {
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

func getLine(r io.Reader) chan Result {
	scanner := bufio.NewScanner(r)
	out := make(chan Result)

	go func() {
		defer close(out)
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			out <- Result{words, nil}
		}

		if err := scanner.Err(); err != nil {
			out <- Result{nil, err}
			return
		}
	}()

	return out
}

func withChan(filename string) ([]int, error) {
	defer timeFunc(time.Now(), "W")
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	funcs := []CalculatorFunc{calculateRound, calculateStrategy}

	sums := make([]int, len(funcs))
	for res := range getLine(file) {

		if res.Err != nil {
			return nil, res.Err
		}

		for i, f := range funcs {
			pts, err := f(res.Words)
			if err != nil {
				return nil, err
			}
			sums[i] += pts
		}
	}
	return sums, nil

}

func withoutChan(filename string) ([]int, error){
	defer timeFunc(time.Now(), "W/o")
	first, err := calculateSum(filename, calculateRound)
	if err != nil {
		return nil, err
	}

	second, err := calculateSum(filename, calculateStrategy)
	if err != nil {
		return nil, err
	}

	return []int{first, second}, nil
}

func main() {
	filename := "input.txt"
	first, err := withChan(filename)
	if err != nil {
		log.Panic(err)
	}

	second, err := withoutChan(filename)
	if err != nil {
		log.Panic(err)
	}
	log.Println(first, second)
}
