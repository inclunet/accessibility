package accessibility

import (
	"github.com/PuerkitoBio/goquery"
)

type Images struct {
	Element
}

func (i *Images) AlternativeDescription() (string, bool) {
	if value, ok := i.Selection.Attr("alt"); ok {
		if value != "" {
			return value, true
		}
	}
	return "", false
}

func (i *Images) isValidAlternativeDescription() bool {
	alt, ok := i.AlternativeDescription()

	if ok {
		if len(alt) >= 3 {
			return true
		}
	}
	return false
}

func (i *Images) Check() (int, bool, string) {
	description := i.isValidAlternativeDescription()
	hidden := i.AriaHidden()

	if !description && !hidden {
		return 1, false, "No hidden imagens needs a descriptionfor accessibility"
	}

	return 1, true, "There is no errors on your image alternative text description."
}

func NewImageCheck(s *goquery.Selection) (int, bool, string) {
	image := new(Images)
	image.Selection = s
	A, Pass, Description := image.Check()
	return A, Pass, Description
}
