package uule

import (
	"bufio"
	"compress/gzip"
	"embed"
	"math"
)

//go:embed canonical-names.gz
var canonicalNames embed.FS

func getCanonicalNames() ([]string, error) {
	// Open the compressed file
	inputFile, err := canonicalNames.Open("canonical-names.gz")
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	// Create a gzip reader
	gr, err := gzip.NewReader(inputFile)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	// Read the decompressed data line by line
	scanner := bufio.NewScanner(gr)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func getClosestMatch(sortedList []string, target string) string {
	n := len(sortedList)
	if n == 0 {
		return ""
	}

	// Binary search to find the closest index
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if sortedList[mid] == target {
			return sortedList[mid]
		} else if sortedList[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Check the boundaries
	if left >= n {
		return sortedList[n-1]
	}
	if right < 0 {
		return sortedList[0]
	}

	// Compare the closest two strings
	closestIndex := left
	if left > 0 && math.Abs(float64(len(target)-len(sortedList[left-1]))) < math.Abs(float64(len(target)-len(sortedList[left]))) {
		closestIndex = left - 1
	}

	return sortedList[closestIndex]
}

func (uule *UULEParams) LookupCanonicalName() *UULEParams {
	// Read the compressed file
	lines, err := getCanonicalNames()
	if err != nil {
		panic(err)
	}

	// Binary search to find the closest match
	closestMatch := getClosestMatch(lines, uule.City+","+uule.Region+","+uule.Country)
	if closestMatch == "" {
		panic("No closest match found")
	}

	// Parse the closest match
	uule.CanonicalName = closestMatch

	return uule
}
