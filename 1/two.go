package main

func replaceEntry(maxIdx, max []int, newMax int, posInner, posOuter int) ([]int, []int) {

	temp := max[posInner]
	max[posInner] = newMax

	tempIdx := maxIdx[posInner]
	maxIdx[posInner] = posOuter

	if posInner >= len(max)-1 {
		return maxIdx, max
	}

	for i := range max[posInner+1 : len(max)-1] {

		max[posInner+i+1] = temp
		temp = max[posInner+i+2]

		maxIdx[posInner+i+1] = tempIdx
		tempIdx = maxIdx[posInner+i+2]

	}

	maxIdx[len(max)-1] = tempIdx
	max[len(max)-1] = temp
	return maxIdx, max
}

func findNMax(elfData [][]int, nMax int) ([]int, []int) {

	if nMax < 0 {
		return nil, nil
	}

	maxNIdx := make([]int, nMax)
	maxN := make([]int, nMax)

	for i, elf := range elfData {
		cals := sumCalories(elf)
		for k, entry := range maxN {
			if cals > entry {
				maxNIdx, maxN = replaceEntry(maxNIdx, maxN, cals, k, i)
				break
			}
		}
	}

	return maxNIdx, maxN
}
