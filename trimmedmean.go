package trimmedmean

import (
	"errors"
	"sort"
)

// TrimmedMean calculates the trimmed mean of a slice of float64s.
// If only one trim proportion is provided, symmetric trimming is assumed.
func TrimmedMean(data []float64, trim ...float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("empty data slice")
	}

	sort.Float64s(data)
	n := len(data)

	var lowerTrim, upperTrim float64

	if len(trim) == 1 {
		// symmetric trimming
		lowerTrim = trim[0]
		upperTrim = trim[0]
	} else if len(trim) == 2 {
		lowerTrim = trim[0]
		upperTrim = trim[1]
	} else {
		return 0, errors.New("provide either one or two trimming values")
	}

	start := int(float64(n) * lowerTrim)
	end := n - int(float64(n)*upperTrim)

	if start >= end {
		return 0, errors.New("trimming proportions are too large, resulting in no data")
	}

	sum := 0.0
	for _, v := range data[start:end] {
		sum += v
	}
	return sum / float64(end-start), nil
}
