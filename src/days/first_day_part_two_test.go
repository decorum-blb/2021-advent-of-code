package main

import (
	"github.com/decorum-blb/2021-advent-of-code/readfile"
	"testing"
)

func TestWindowSizeOfZero(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{1, 2, 3}, 0)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestEmptyArray(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{}, 3)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestInputSmallerThanWindow(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{1, 2}, 3)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestNotEnoughGroupsToCompareWithGivenWindow(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{1, 2, 3}, 3)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestOneGroupIsLargerThanPrevious(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{1, 2, 3, 4}, 3)

	if actual != 1 {
		t.Errorf("Actual = %d; Should have been 1", actual)
	}
}

func TestNoGroupIsLargerThanPrevious(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 3)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestExampleInput(t *testing.T) {
	actual := CountLargerMeasurementsByWindow([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 3)

	if actual != 5 {
		t.Errorf("Actual = %d; Should have been 5", actual)
	}
}

func TestAdventOfCodeTestDataForPartTwo(t *testing.T) {
	data, err := readfile.ReadFile("./first_dat_test_data.txt")

	if err != nil {
		panic(err)
	}

	actual := CountLargerMeasurementsByWindow(data, 3)

	if actual != 1139 {
		t.Errorf("Actual = %d; Should have been 1139", actual)
	}
}
