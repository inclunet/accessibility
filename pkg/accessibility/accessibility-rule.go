package accessibility

import (
	"encoding/json"
	"os"
)

type AccessibilityRule struct {
	A           int
	Description string
	Error       bool
	Solution    string
	Warning     bool
}

func LoadAccessibilityRules(filename string) (map[string]AccessibilityRule, error) {
	accessibilityRules := make(map[string]AccessibilityRule)
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
