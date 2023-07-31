package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type Result struct {
	Line string
	Err error
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}


func byteToPriority(b byte) (int, error) {
	if b >= 97 && b <= 122 {
		return int(b - 96), nil
	}

	if b >= 65 && b <= 90 {
		return int(b - 38), nil
	}

	return 0, errors.New("incorrect item")
}

func selectCommon(line string) (byte, error) {
	byteArray := []byte(line)

	firstHalf := byteArray[:len(byteArray) / 2]
	secondHalf := byteArray[len(byteArray) / 2:]

	symset := make(map[byte]bool)
	for _, b := range firstHalf {
		if !symset[b] {
			symset[b] = true
		}
	}

	for _, b := range secondHalf {
		if symset[b] {
			return b, nil
		}
	}
	return 0, errors.New("no common items.")
}

func getLine(filename string) (chan *Result, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	out := make(chan *Result)

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(out)
		for scanner.Scan() {
			out <- &Result{scanner.Text(), nil}
		}

		if err := scanner.Err(); err != nil {
			out <- &Result{"", err}
			return
		}
	}()

	return out, nil
}

func one(filename string) (int, error) {
	out, err := getLine(`input.txt`)
	if err != nil {
		return 0, err
	}

	sum := 0
	for l := range out {
		if err := l.Err; err != nil {
			return 0, err
		}

		item, err := selectCommon(l.Line)
		if err != nil {
			return 0, err
		}

		p, err := byteToPriority(item)

		if err != nil {
			return 0, err
		}
		sum += p
	}

	return sum, nil

}
	

func main() {
	filename := "input.txt"
	res, err := one(filename)
	panicOnErr(err)
	log.Println(res)

}