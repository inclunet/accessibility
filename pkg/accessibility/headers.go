package accessibility

import (
	"strconv"
	"strings"
)

type Headers struct {
	Element
}

func (h *Headers) GetHeaderLevel(accessibilityCheck AccessibilityCheck) (int, bool) {
	if h.isHeader(accessibilityCheck) {
		level, _ := strconv.Atoi(strings.TrimPrefix(accessibilityCheck.Element, "h"))
		return level, true
	}

	return 0, false
}

func (h *Headers) isHeader(accessibilityCheck AccessibilityCheck) bool {
	switch accessibilityCheck.Element {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}

func (h *Headers) isIncorrectLevel(headerLevel int, currentLevel int) bool {
	return currentLevel > headerLevel+1
}

func (h *Headers) CheckHierarchy(accessibilityCheck AccessibilityCheck) bool {
	accessibilityChecks := append(h.AccessibilityChecks, accessibilityCheck)
	headerLevel := 0

	for _, accessibilityCheck := range accessibilityChecks {
		if currentLevel, ok := h.GetHeaderLevel(accessibilityCheck); ok {
			if h.isIncorrectLevel(headerLevel, currentLevel) {
				return true
			}

			headerLevel = currentLevel
		}
	}

	return false
}

func (h *Headers) Check() AccessibilityCheck {
	accessibilityCheck := h.NewAccessibilityCheck("pass")

	if h.AriaHidden() {
		return h.FindViolation(accessibilityCheck, "aria-hidden")
	}

	if h.CheckHierarchy(accessibilityCheck) {
		return h.FindViolation(accessibilityCheck, "emag-1.3.2")
	}

	accessibleText, ok := h.AccessibleText()

	if !ok {
		return h.FindViolation(accessibilityCheck, "emag-1.2.3")
	}

	if h.CheckTooShortText(accessibleText) {
		return h.FindViolation(accessibilityCheck, "too-short-text")
	}

	if h.CheckTooLongText(accessibleText, 80) {
		return h.FindViolation(accessibilityCheck, "too-long-text")
	}

	return accessibilityCheck
}
