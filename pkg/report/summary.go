package report

import (
	"math"

	"github.com/inclunet/accessibility/pkg/accessibility"
)

type AccessibilitySummary struct {
	Checks   []accessibility.AccessibilityCheck
	Errors   int
	Pass     int
	Rat      float32
	Total    int
	Warnings int
}

func (s *AccessibilitySummary) Update(accessibilityCheck accessibility.AccessibilityCheck) {
	s.Total++

	if accessibilityCheck.Warning {
		s.Warnings++
	}

	if accessibilityCheck.Error {
		s.Pass++
	} else {
		s.Errors++
	}

	s.UpdateRat()
	s.Checks = append(s.Checks, accessibilityCheck)
}

func (s *AccessibilitySummary) UpdateRat() {
	if s.Errors > 0 {
		rat := (100.0 * 10 / float32(s.Total*10)) * float32(s.Pass*10)
		s.Rat = float32(math.Ceil(float64(rat / 10)))
	}
}

func NewSummary() *AccessibilitySummary {
	Summary := new(AccessibilitySummary)
	return Summary
}
