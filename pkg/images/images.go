package images

import (
	"github.com/PuerkitoBio/goquery"
)

func isAriaHidden(s *goquery.Selection) bool {
	value, exists := s.Attr("aria-hidden")

	if exists == false {
		return false
	}

	if value == "true" {
		return true
	} else {
		return false
	}
}

func isValidAlternativeDescription(s *goquery.Selection) bool {
	value, exists := s.Attr("alt")

	if exists == false {
		return false
	}

	if value == "" {
		return false
	}

	if len(value) <= 3 {
		return false
	}

	return true
}

func Check(s *goquery.Selection) (int, bool, string) {
	description := isValidAlternativeDescription(s)
	hidden := isAriaHidden(s)

	if description == false && hidden == false {
		return 1, false, "No hidden imagens needs a descriptionfor accessibility"
	}

	return 1, true, "There is no errors on your image alternative text description."
}
