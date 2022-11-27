package accessibility

type Accessibility interface {
	AlternativeText() (string, bool)
	AriaHidden() bool
	AriaLabel() (string, bool)
	Check() (int, bool, string)
	Role() (string, bool)
}
