package trimmedmean

import (
	"math"
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 100}
	mean, err := TrimmedMean(data, 0.2)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expected := 3.5
	if math.Abs(mean-expected) > 1e-6 {
		t.Errorf("unexpected result: got %v, expected %v", mean, expected)
	}
}
