// Package pages contains all page component definitions for the application.
// This file defines the Contact page and its form component.
package pages

import (
	"form_exer/web/shared" // Local package with shared components
	"github.com/rohanthewiz/element" // Third-party HTML builder library
)

// STRUCT DEFINITION with EMBEDDING
// ContactPage represents the contact page structure
// This is nearly identical to the Home page structure - demonstrating code reuse
type ContactPage struct {
	// EMBEDDED FIELD: shared.Page provides Banner() and Footer() methods
	// This is the MIXIN PATTERN - ContactPage inherits Page's functionality
	shared.Page

	// Page-specific field for the heading
	Heading string
}

// PACKAGE-LEVEL VARIABLE: Contact is a singleton instance
// This is created at package initialization time (before main() runs)
// EXPORTED: Capital 'C' makes it accessible from other packages
var Contact = ContactPage{
	// NESTED STRUCT LITERAL: Initializing the embedded Page field
	Page: shared.Page{Title: "Contact Us"},

	// Initialize the Heading field specific to this page
	Heading: "Get in Touch",
}

// METHOD with VALUE RECEIVER and NAMED RETURN VALUE
// (c ContactPage) - value receiver, this is a method on ContactPage type
// (out string) - NAMED RETURN VALUE (declared but not explicitly used)
// This method is almost identical to Home.Render() - showing consistent patterns
func (c ContactPage) Render() (out string) {
	// Create a new HTML builder for this page
	b := element.NewBuilder()

	// METHOD CHAINING: Build the page structure
	// The pattern is: body → components (banner, form, footer) → heading
	b.Body("style", "background-color:tan").R(
		// COMPOSITE PATTERN: Render multiple components together
		// element.RenderComponents takes a builder and multiple components
		element.RenderComponents(b,
			// METHOD from EMBEDDED FIELD: c.Banner() works due to embedding
			// Equivalent to c.Page.Banner() but Go allows the shorthand
			c.Banner(), // Renders the page banner at the top

			// EMPTY STRUCT LITERAL: Creating ContactForm instance inline
			// ContactForm{} creates a zero-value instance
			ContactForm{}, // Renders the contact form

			// Another method from the embedded Page
			c.Footer(), // Renders the page footer at the bottom
		),
		// Add the page heading after the components
		// c.Heading accesses the ContactPage's Heading field
		b.H1("style", "color:maroon;background-color:#dfc673").T(c.Heading),
	)

	// Convert the builder to an HTML string and return it
	return b.String()
}

// EMPTY STRUCT for FORM COMPONENT
// ContactForm is a stateless component - no data fields needed
// All form structure and attributes are defined in the Render method
type ContactForm struct{}

// METHOD with POINTER PARAMETER and NAMED RETURN
// (cf ContactForm) - value receiver for the empty struct
// (b *element.Builder) - POINTER parameter to avoid copying the builder
// (dontCare any) - named return with 'any' type (we return nil via naked return)
func (cf ContactForm) Render(b *element.Builder) (dontCare any) {
	// FORM ELEMENT: Build an HTML form with attributes
	// ATTRIBUTE PAIRS: "name", "value", "name", "value" pattern
	// action="/contact" - where to send form data (POST request to /contact endpoint)
	// method="POST" - HTTP method for form submission (POST for data modification)
	b.Form("action", "/contact", "method", "POST").R(
		// INPUT ELEMENT: Text input field
		// MULTIPLE ATTRIBUTES demonstrated:
		//   type="text" - standard text input (single line)
		//   name="name" - field name used when submitting form data
		//   placeholder="Name" - hint text shown when field is empty
		b.Input("type", "text", "name", "name", "placeholder", "Name"),

		// INPUT ELEMENT: Email input field
		// type="email" - HTML5 input type that validates email format
		// Browser will enforce basic email validation before submission
		b.Input("type", "email", "name", "email", "placeholder", "Email"),

		// TEXTAREA ELEMENT: Multi-line text input
		// name="message" - field identifier for form submission
		// .R() with no arguments creates an empty textarea (no child elements)
		// TextArea is different from Input - it's a paired tag (<textarea></textarea>)
		b.TextArea("name", "message", "placeholder", "Message").R(),

		// BUTTON ELEMENT: Submit button
		// type="submit" - clicking this button submits the form
		// .T("Send") adds text content to the button
		// When clicked, browser sends POST request to /contact with form data
		b.Button("type", "submit").T("Send"),
	)

	// NAKED RETURN: Returns the named value 'dontCare' (which is nil by default)
	// We don't need to return anything meaningful, so we use a naked return
	return
}

// KEY CONCEPTS demonstrated in this file:
// 1. STRUCT EMBEDDING - ContactPage embeds shared.Page (mixin pattern)
// 2. PACKAGE-LEVEL VARIABLES - Contact singleton created at init time
// 3. NAMED RETURN VALUES - Enable naked returns and self-documentation
// 4. VALUE RECEIVERS - Methods receive copies of structs
// 5. POINTER PARAMETERS - Avoid copying large structs (builder)
// 6. EMPTY STRUCTS - Zero-size structs for stateless components
// 7. FORM ELEMENTS - Input, TextArea, Button with proper attributes
// 8. HTML5 INPUT TYPES - email type with built-in validation
// 9. FORM SUBMISSION - POST method to server endpoint
// 10. CONSISTENT PATTERNS - Similar structure to Home page for maintainability
