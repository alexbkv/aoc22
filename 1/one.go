package main

func sumCalories(elf []int) int {
	sum := 0
	for _, cal := range elf {
		sum += cal
	}
	return sum
}

func findMaxCalories(elfData [][]int) (int, int) {
	maxIdx := 0
	maxCalories := 0

	for i, elf := range elfData {
		calories := sumCalories(elf)
		if calories > maxCalories {
			maxIdx = i
			maxCalories = calories
		}
	}

	return maxIdx, maxCalories
}
