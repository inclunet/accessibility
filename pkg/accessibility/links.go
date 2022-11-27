package accessibility

import "github.com/PuerkitoBio/goquery"

type Links struct {
	Element
}

func (l *Links) Check() (int, bool, string) {
	if !l.AriaHidden() {
		if accessibleText, ok := l.AccessibleText(); ok && len(accessibleText) > 3 {
			return 1, true, "This link are providing a valid    description text for screen readers."
		}
		return 1, false, "If your link is not hidden, you need a text description for screen reader software."
	}
	return 1, true, "Hidden Links do not need text description."
}

func NewLinkCheck(s *goquery.Selection) Accessibility {
	accessibilityInterface := new(Links)
	accessibilityInterface.Selection = s
	return accessibilityInterface
}
