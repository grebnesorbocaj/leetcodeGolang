package main

import "fmt"

func main() {
	numRows := 5
	triangle := [][]int{}
	for currentRow := 0; currentRow < numRows; currentRow++ {
		currentRowSlice := []int{}
		for currentIndex := 0; currentIndex <= currentRow; currentIndex++ {
			posVal := findVal(currentRow, currentIndex, triangle)
			currentRowSlice = append(currentRowSlice, posVal)
		}
		triangle = append(triangle, currentRowSlice)
	}

	fmt.Println(triangle)
}

func findVal(currentRow int, currentIndex int, currentTriangle [][]int) int {
	var value int = 0
	if currentIndex > currentRow-1 {
		return 1
	}
	for currentIndex >= 0 {
		value += currentTriangle[currentRow-1][currentIndex]
		currentIndex--
		currentRow--
	}
	return value
}
