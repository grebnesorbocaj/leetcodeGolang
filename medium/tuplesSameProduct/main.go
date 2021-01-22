package main

func main() {

}

func finalCalc(t []int) int {
	if len(t) < 4 {
		return 0
	}
	return calcTotalSets(t) * 8
}

func calcTotalSets(tuple []int) int {
	endOf := len(tuple) - 1
	comboMap := make(map[int]int)

	for i := range tuple {
		for j := i + 1; j <= endOf; j++ {
			comboMap[tuple[i]*tuple[j]] += 1
		}
	}

	totalSets := 0
	for _, num := range comboMap {
		totalSets += calculateNumCombos(num)
	}

	return totalSets
}

func calculateNumCombos(x int) int {
	var sum int = 0
	for x >= 1 {
		x--
		sum += x
	}
	return sum
}
