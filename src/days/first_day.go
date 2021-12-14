package main

func CountLargerMeasurements(scans []int) int {
	count := 0

	for index, scan := range scans {
		if index > 0 && scans[index-1] < scan {
			count += 1
		}
	}

	return count
}

func CountLargerMeasurementsByWindow(scans []int, windowSize int) int {
	if windowSize < 1 || len(scans) < windowSize {
		return 0
	} else {
		return countStuff(scans, windowSize, 0, 0)
	}
}

func countStuff(baseArray []int, windowSize int, counter int, index int) int {
	if index+1+windowSize > len(baseArray) {
		return counter
	}
	groupA := sum(getGroup(baseArray, windowSize, index))
	groupB := sum(getGroup(baseArray, windowSize, index+1))

	newCounter := counter
	if groupB > groupA {
		newCounter += 1
	}

	return countStuff(baseArray, windowSize, newCounter, index+1)
}

func getGroup(baseArray []int, windowSize int, index int) []int {
	var group []int = []int{}

	for i := 0; i < windowSize; i++ {
		elem := baseArray[index+i]
		group = append(group, elem)
	}

	return group
}

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
