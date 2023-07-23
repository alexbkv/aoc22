package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func sumCalories(elf []int) int {
	sum := 0
	for _, cal := range elf {
		sum += cal
	}
	return sum
}

func readFile(name string) ([][]int, error) {

	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]int
	var elf []int

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			result = append(result, elf)
			elf = nil
			continue
		}
		trimmed := strings.Trim(text, "\n")
		num, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, err
		}
		elf = append(elf, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
