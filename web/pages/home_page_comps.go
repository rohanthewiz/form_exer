package pages

import (
	"form_exer/web/shared"
	"github.com/rohanthewiz/element"
)

// Home Page component
type Home struct {
	shared.Page // Mixin for common page components
	Heading     string
}

func (h Home) Render() (out string) {
	b := element.NewBuilder()

	b.Body("style", "background-color:tan").R(
		element.RenderComponents(b,
			h.Banner(),      // Common banner from Page mixin
			ContactForm{},   // Page-specific content
			h.Footer(),      // Common footer from Page mixin
		),
		b.H1("style", "color:maroon;background-color:#dfc673").T(h.Heading),
	)

	return b.String()
}

// ContactForm component renders a contact form using element package
// TODO: Put contact form and footer logic in separate files
type ContactForm struct{}

func (c ContactForm) Render(b *element.Builder) (dontCare any) {
	// Build the contact form with proper attributes and structure
	b.Form("action", "/contact", "method", "POST").R(
		b.Input("type", "text", "name", "name", "placeholder", "Name"),
		b.Input("type", "email", "name", "email", "placeholder", "Email"),
		b.TextArea("name", "message", "placeholder", "Message").R(),
		b.Button("type", "submit").T("Send"),
	)

	return
}

