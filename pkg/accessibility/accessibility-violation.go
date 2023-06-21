package accessibility

import (
	"encoding/json"
	"os"
)

type AccessibilityViolation struct {
	A           int
	Description string
	Error       bool
	Solution    string
	Warning     bool
}

func LoadAccessibilityViolations(filename string) (map[string]AccessibilityViolation, error) {
	accessibilityRules := make(map[string]AccessibilityViolation)
	accessibilityRulesFile, err := os.ReadFile(filename)

	if err != nil {
		return accessibilityRules, err
	}

	err = json.Unmarshal(accessibilityRulesFile, &accessibilityRules)

	if err != nil {
		return accessibilityRules, err
	}

	return accessibilityRules, nil
}
