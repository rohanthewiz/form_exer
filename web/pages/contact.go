package pages

import (
	"form_exer/web/shared"
	"github.com/rohanthewiz/element"
)

// ContactPage component
type ContactPage struct {
	shared.Page // Mixin for common page components
	Heading     string
}

var Contact = ContactPage{
	Page:    shared.Page{Title: "Contact Us"},
	Heading: "Get in Touch",
}

func (c ContactPage) Render() (out string) {
	b := element.NewBuilder()

	b.Body("style", "background-color:tan").R(
		element.RenderComponents(b,
			c.Banner(),      // Common banner from Page mixin
			ContactForm{},   // Contact form component
			c.Footer(),      // Common footer from Page mixin
		),
		b.H1("style", "color:maroon;background-color:#dfc673").T(c.Heading),
	)

	return b.String()
}

// ContactForm component renders a contact form using element package
type ContactForm struct{}

func (cf ContactForm) Render(b *element.Builder) (dontCare any) {
	// Build the contact form with proper attributes and structure
	b.Form("action", "/contact", "method", "POST").R(
		b.Input("type", "text", "name", "name", "placeholder", "Name"),
		b.Input("type", "email", "name", "email", "placeholder", "Email"),
		b.TextArea("name", "message", "placeholder", "Message").R(),
		b.Button("type", "submit").T("Send"),
	)

	return
}
