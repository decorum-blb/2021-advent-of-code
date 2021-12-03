package main

import "testing"

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
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	actual := CountLargerMeasurements(data)

	if actual != 7 {
		t.Errorf("Actual = %d; Should have been 7", actual)
	}
}
