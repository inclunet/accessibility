package accessibility

import (
	"strconv"
	"strings"
)

type Headers struct {
	Element
}

func (h *Headers) Check() AccessibilityCheck {
	accessibilityCheck := h.NewAccessibilityCheck("pass")

	if h.AriaHidden() {
		return accessibilityCheck.SetViolation("aria-hidden")
	}

	if HeaderWithoutMain(h.AccessibilityChecks, accessibilityCheck) {
		return accessibilityCheck.SetViolation("emag-1.3.5")
	}

	if HeaderOverflow(h.AccessibilityChecks, accessibilityCheck) {
		return accessibilityCheck.SetViolation("emag-1.3.6")
	}

	if HeaderInvalidHierarchy(h.AccessibilityChecks, accessibilityCheck) {
		return accessibilityCheck.SetViolation("emag-1.3.2")
	}

	accessibleText, ok := h.AccessibleText()

	if !ok {
		return accessibilityCheck.SetViolation("emag-1.2.3")
	}

	if h.CheckTooShortText(accessibleText) {
		return accessibilityCheck.SetViolation("too-short-text")
	}

	if h.CheckTooLongText(accessibleText, 80) {
		return accessibilityCheck.SetViolation("too-long-text")
	}

	return accessibilityCheck
}

func HeaderCheck(accessibilityCheck AccessibilityCheck) bool {
	switch accessibilityCheck.Element {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}

func HeaderCount(accessibilityChecks []AccessibilityCheck) int {
	headerCount := 0
	for _, accessibilityCheck := range accessibilityChecks {
		if headerLevel, ok := HeaderLevel(accessibilityCheck); ok && headerLevel >= 1 && headerLevel <= 6 {
			headerCount++
		}
	}
	return headerCount
}

func HeaderIncorrectLevel(headerLevel int, currentLevel int) bool {
	return currentLevel > headerLevel+1
}

func HeaderInvalidHierarchy(accessibilityChecks []AccessibilityCheck, accessibilityCheck AccessibilityCheck) bool {
	accessibilityChecks = append(accessibilityChecks, accessibilityCheck)
	headerLevel := 0

	for _, accessibilityCheck := range accessibilityChecks {
		if currentLevel, ok := HeaderLevel(accessibilityCheck); ok {
			if HeaderIncorrectLevel(headerLevel, currentLevel) {
				return true
			}

			headerLevel = currentLevel
		}
	}

	return false
}

func HeaderInvalidOrdenation(accessibilityChecks []AccessibilityCheck) bool {
	return HeaderMainCount(accessibilityChecks) == 0 && HeaderCount(accessibilityChecks) > 0
}

func HeaderLevel(accessibilityCheck AccessibilityCheck) (int, bool) {
	if HeaderCheck(accessibilityCheck) {
		level, _ := strconv.Atoi(strings.TrimPrefix(accessibilityCheck.Element, "h"))
		return level, true
	}
	return 0, false
}

func HeaderMainCount(accessibilityChecks []AccessibilityCheck) int {
	headerCount := 0
	for _, accessibilityCheck := range accessibilityChecks {
		if headerLevel, ok := HeaderLevel(accessibilityCheck); ok && headerLevel == 1 {
			headerCount++
		}
	}
	return headerCount
}

func HeaderMainUnavailable(accessibilityChecks []AccessibilityCheck) bool {
	return HeaderMainCount(accessibilityChecks) == 0
}

func HeaderOverflow(accessibilityChecks []AccessibilityCheck, accessibilityCheck AccessibilityCheck) bool {
	return HeaderMainCount(append(accessibilityChecks, accessibilityCheck)) > 1
}

func HeaderUnavailable(accessibilityChecks []AccessibilityCheck) bool {
	return HeaderCount(accessibilityChecks) == 0
}

func HeaderWithoutMain(accessibilityChecks []AccessibilityCheck, accessibilityCheck AccessibilityCheck) bool {
	headerCount := 0
	accessibilityChecks = append(accessibilityChecks, accessibilityCheck)
	for _, check := range accessibilityChecks {
		if headerLevel, ok := HeaderLevel(check); ok {
			if headerLevel >= 2 && headerLevel <= 6 && headerCount == 0 {
				return true
			}

			headerCount++
		}
	}

	return false
}
