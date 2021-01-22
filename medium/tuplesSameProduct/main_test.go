package main

import "testing"

func TestCalcTotalSets(t *testing.T) {
	if calcTotalSets([]int{2, 3, 4, 6})*8 != 8 {
		t.Errorf("Expected response: 8, but got: %v", calcTotalSets([]int{2, 3, 4, 6, 8, 12})*8)
	}

	if calcTotalSets([]int{1, 2, 4, 5, 10})*8 != 16 {
		t.Errorf("Expected response: 16, but got: %v", calcTotalSets([]int{2, 3, 4, 6, 8, 12})*8)
	}

	if calcTotalSets([]int{2, 3, 4, 6, 8, 12})*8 != 40 {
		t.Errorf("Expected response: 40, but got: %v", calcTotalSets([]int{2, 3, 4, 6, 8, 12})*8)
	}

	if calcTotalSets([]int{2, 3, 5, 7})*8 != 0 {
		t.Errorf("Expected response: 0, but got: %v", calcTotalSets([]int{2, 3, 4, 6, 8, 12})*8)
	}
}

func TestFinalCalc(t *testing.T) {
	if finalCalc([]int{2, 3, 4, 6}) != 8 {
		t.Errorf("Expected response: 8, but got: %v", finalCalc([]int{2, 3, 4, 6, 8, 12}))
	}

	if finalCalc([]int{1, 2, 4, 5, 10}) != 16 {
		t.Errorf("Expected response: 16, but got: %v", finalCalc([]int{2, 3, 4, 6, 8, 12}))
	}

	if finalCalc([]int{2, 3, 4, 6, 8, 12}) != 40 {
		t.Errorf("Expected response: 40, but got: %v", finalCalc([]int{2, 3, 4, 6, 8, 12}))
	}

	if finalCalc([]int{2, 3, 5, 7}) != 0 {
		t.Errorf("Expected response: 0, but got: %v", finalCalc([]int{2, 3, 4, 6, 8, 12}))
	}
}
