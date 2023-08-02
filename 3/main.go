package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type Result struct {
	Line string
	Err  error
}

type Set map[int]bool

func (s Set) append(num int) {
	if _, ok := s[num]; !ok {
		s[num] = true
	}
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

func processLine(symset map[byte]Set, line string, lineNum int) {

	lineBytes := []byte(line)
	for _, b := range lineBytes {
		if _, ok := symset[b]; ok {
			symset[b].append(lineNum)
		}
	}
}

func selectCommon(lines ...string) (byte, error) {

	if len(lines) < 2 {
		return 0, errors.New("provide at least 2 lines two find common symbol")
	}

	symset := make(map[byte]Set)

	for _, b := range []byte(lines[0]) {
		_, ok := symset[b]
		if !ok {
			set := make(Set)
			set.append(1)
			symset[b] = set
		}
	}

	for i, line := range lines[1:] {
		processLine(symset, line, i+2)
	}

	for k := range symset {
		if len(symset[k]) == len(lines) {
			return k, nil
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

		size := len(l.Line)
		p, err := linesPriority(l.Line[:size/2], l.Line[size/2:])
		if err != nil {
			return 0, err
		}
		sum += p
	}

	return sum, nil

}

func two(filename string) (int, error) {
	out, err := getLine(`input.txt`)
	if err != nil {
		return 0, err
	}

	sum := 0

	groupSize := 3
	var group []string

	for l := range out {
		if err := l.Err; err != nil {
			return 0, err
		}
		if len(group) == groupSize {
			p, err := linesPriority(group...)
			if err != nil {
				return 0, err
			}
			sum += p
			group = nil
		}
		group = append(group, l.Line)
	}

	p, err := linesPriority(group...)
	if err != nil {
		return 0, err
	}
	sum += p
	group = nil

	return sum, nil
}

func linesPriority(lines ...string) (int, error) {
	item, err := selectCommon(lines...)
	if err != nil {
		return 0, err
	}

	p, err := byteToPriority(item)

	if err != nil {
		return 0, err
	}
	return p, nil
}

func main() {
	filename := "input.txt"
	res, err := one(filename)
	panicOnErr(err)
	log.Println(res)

	res, err = two(filename)
	panicOnErr(err)
	log.Println(res)

}
