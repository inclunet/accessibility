package summary

type AccessibilityCheck struct {
	Element     string
	A           int
	Pass        bool
	Description string
	Line        int
	Html        string
}

type AccessibilitySummary struct {
	Element string
	Pass    int
	Errors  int
	Total   int
}

func GetIndex(Summary []AccessibilitySummary, element string) int {
	for i, entry := range Summary {
		if entry.Element == element {
			return i
		}
	}
	return -1
}

func NewEntry(element string) AccessibilitySummary {
	return AccessibilitySummary{Element: element, Pass: 0, Errors: 0, Total: 0}
}

func GetEntry(Summary []AccessibilitySummary, element string) (int, AccessibilitySummary) {
	i := GetIndex(Summary, element)
	if i >= 0 {
		return i, Summary[i]
	}
	return i, NewEntry(element)
}

func Generate(AccessibilityChecks []AccessibilityCheck) []AccessibilitySummary {
	var Summary []AccessibilitySummary
	for _, check := range AccessibilityChecks {
		i, entry := GetEntry(Summary[:], check.Element)
		entry.Total = entry.Total + 1
		if check.Pass == true {
			entry.Pass = entry.Pass + 1
		} else {
			entry.Errors = entry.Errors + 1
		}
		if i >= 0 {
			Summary[i] = entry
		} else {
			Summary = append(Summary, entry)
		}
	}
	return Summary
}
