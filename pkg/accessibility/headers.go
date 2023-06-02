package accessibility

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
				return false
			}

			headerLevel = currentLevel
		}
	}

	return true
}

func (h *Headers) Check() AccessibilityCheck {
	accessibilityCheck := h.NewAccessibilityCheck(1, "Please check if this header is following the headers hierarchy")

	if h.AriaHidden() {
		accessibilityCheck.Pass = true
		accessibilityCheck.Description = "Aria-hidden headers do not need to follow ierarchi or needs a accessibility text description"
		return accessibilityCheck
	}

	if h.CheckHierarchy(accessibilityCheck) {
		accessibilityCheck.Description = "This header is following correct ierarchi but do not have a text description"

		if accessibleText, ok := h.AccessibleText(); ok {
			accessibilityCheck.Pass = true
			accessibilityCheck.Description = "This header are ok"

			if h.CheckTooShortText(accessibleText) {
				accessibilityCheck.Description = "Small headers can not provide information to screen readers"
			}

			if h.CheckTooLongText(accessibleText, 80) {
				accessibilityCheck.Description = "Please check if this header is too longh, too long headers is not presented at a single line and is possibly a bad practice."
			}
		}
	}

	return accessibilityCheck
}

func NewHeaderCheck(s *goquery.Selection, accessibilityChecks []AccessibilityCheck) Accessibility {
	accessibilityInterface := new(Headers)
	accessibilityInterface.Selection = s
	accessibilityInterface.AccessibilityChecks = accessibilityChecks
	return accessibilityInterface
}
