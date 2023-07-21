package main

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
