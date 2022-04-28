package levenshtein

import (
	"fmt"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	s1, s2 := "kitten", "sitting"
	distance := Calculate(s1, s2)
	if distance != 3 {
		t.Errorf("The distance of %s and %s should be 3, not %d\n", s1, s2, distance)
	}
}

func TestCalculatePercentage(t *testing.T) {
	s1, s2 := "abc", "ab"
	percentage := CalculatePercentage(s1, s2)
	fmt.Println(percentage)
}
