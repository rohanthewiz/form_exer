package shared

// Page is a mixin struct that provides common page components
type Page struct {
	Title string
}

// Banner returns the banner component for this page
func (p Page) Banner() Banner {
	return Banner{Title: p.Title}
}

// Footer returns the footer component for this page
func (p Page) Footer() Footer {
	return Footer{}
}
