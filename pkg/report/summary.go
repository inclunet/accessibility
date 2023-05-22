package report

import "math"

type AccessibilitySummary struct {
	Pass   int
	Errors int
	Total  int
	Rat    float32
}

func (s *AccessibilitySummary) Count() {
	s.Total = s.Total + 1
	s.UpdateRat()
}

func (s *AccessibilitySummary) AddError() {
	s.Errors = s.Errors + 1
	s.Count()
}

func (s *AccessibilitySummary) AddPass() {
	s.Pass = s.Pass + 1
	s.Count()
}

func (s *AccessibilitySummary) Update(Pass bool) {
	if Pass {
		s.AddPass()
	} else {
		s.AddError()
	}
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
