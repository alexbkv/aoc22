package main

import (
	"log"
)

func main() {
	elfData, err := readFile("input.txt")
	panicOnErr(err)

	_, maxCalories := findMaxCalories(elfData)
	log.Println(maxCalories)
}
