package report

type AccessibilitySummary struct {
	Pass   int
	Errors int
	Total  int
}

func (s *AccessibilitySummary) Count() {
	s.Total = s.Total + 1
}

func (s *AccessibilitySummary) AddError() {
	s.Errors = s.Errors + 1
	s.Count()
}

func (s *AccessibilitySummary) AddPass() {
	s.Pass = s.Pass + 1
	s.Count()
}

func NewSummary() *AccessibilitySummary {
	Summary := new(AccessibilitySummary)
	return Summary
}
