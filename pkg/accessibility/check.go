package accessibility

import "github.com/PuerkitoBio/goquery"

func NewCheck(s *goquery.Selection) (int, bool, string) {
	accessibilityCheckList := map[string]Accessibility{
		"a":   Links,
		"img": Images,
	}
	elementName := goquery.NodeName(s)
	if elementInterface, ok := accessibilityCheckList[elementName]; ok {
		a := new(elementInterface)
	}
	return 1, false, ""
}
