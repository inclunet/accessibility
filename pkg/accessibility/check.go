package accessibility

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

func NewCheck(s *goquery.Selection) (int, bool, string, error) {
	accessibilityCheckList := map[string]func(*goquery.Selection) Accessibility{
		"a":   NewLinkCheck,
		"img": NewImageCheck,
	}
	elementName := goquery.NodeName(s)
	if elementInterface, ok := accessibilityCheckList[elementName]; ok {
		accessibilityInterface := elementInterface(s)
		a, pass, description := accessibilityInterface.Check()
		return a, pass, description, nil
	}
	return 1, false, "", errors.New("No defined evaluator function for this html element")
}
