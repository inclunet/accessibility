package accessibility

import "fmt"

type AmpImg struct {
	Images
}

func (a *AmpImg) Check() ([]AccessibilityCheck, bool) {
	fmt.Println("a")
	if accessibilityChecks, pass := a.Images.Check(); pass {
		fmt.Println("*")
		return accessibilityChecks, pass
	}

	accessibilityChecks := Check(a.Selection.Children(), []AccessibilityCheck{})
	fmt.Println("*")
	for _, accessibilityCheck := range accessibilityChecks {
		fmt.Println("*")
		if accessibilityCheck.Violation != "pass" {
			return a.AddViolation(accessibilityCheck.Violation, false)
		}
	}

	return a.AddViolation("pass", true)
}
