package main

import (
	"github.com/decorum-blb/2021-advent-of-code/readfile"
	"testing"
)
func TestEmptyInput(t *testing.T) {
	actual := CountLargerMeasurements([]int{})

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestSingleElementArray(t *testing.T) {
	data := []int{1}
	actual := CountLargerMeasurements(data)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestTwoElementsWhereSecondIsSmaller(t *testing.T) {
	data := []int{4, 1}
	actual := CountLargerMeasurements(data)

	if actual > 0 {
		t.Errorf("Actual = %d; Should have been 0", actual)
	}
}

func TestTwoElementsWhereSecondIsGreater(t *testing.T) {
	data := []int{1, 4}
	actual := CountLargerMeasurements(data)

	if actual != 1 {
		t.Errorf("Actual = %d; Should have been 1", actual)
	}
}

func TestAdventOfCodeTestData(t *testing.T) {
	data, err := readfile.ReadFile("./first_day_test_data.txt")

	if err != nil {
		panic(err)
	}

	actual := CountLargerMeasurements(data)

	if actual != 1139 {
		t.Errorf("Actual = %d; Should have been 1139", actual)
	}
}
