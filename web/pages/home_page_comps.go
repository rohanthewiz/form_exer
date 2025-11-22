package pages

import "github.com/rohanthewiz/element"

// Home Page component
type Home struct{}

func (h Home) Render() (out string) {
	b := element.NewBuilder()

	var contactForm = ContactForm{}

	b.Body("style", "background-color:tan").R(
		b.H1("style", "color:maroon;background-color:#dfc673").T("Welcome"),
		contactForm.Render(b),
	)

	return b.String()
}

// ContactForm component renders a contact form using element package
type ContactForm struct{}

func (c ContactForm) Render(b *element.Builder) string {
	// Build the contact form with proper attributes and structure
	b.Form("action", "/contact", "method", "POST").R(
		b.Input("type", "text", "name", "name", "placeholder", "Name"),
		b.Input("type", "email", "name", "email", "placeholder", "Email"),
		b.TextArea("name", "message", "placeholder", "Message").R(),
		b.Button("type", "submit").T("Send"),
	)

	return b.String()
}
