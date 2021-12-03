package main

import (
	"io/ioutil"
	"strconv"
	"strings"
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
	data, err := readFile("./dayone_test_data.txt")

	if err != nil {
		panic(err)
	}

	actual := CountLargerMeasurements(data)

	if actual != 1139 {
		t.Errorf("Actual = %d; Should have been 1139", actual)
	}
}

// It would be better for such a function to return error, instead of handling
// it on their own.
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
