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
