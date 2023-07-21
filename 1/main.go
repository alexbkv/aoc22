package main

import (
	"log"
)

func main() {
	elfData, err := readFile("input.txt")
	panicOnErr(err)

	_, maxCalories := findMaxCalories(elfData)
	maxIdx, max := findNMax(elfData, 3)
	log.Println(maxCalories, maxIdx, max)
}
