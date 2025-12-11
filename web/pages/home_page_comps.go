// Package pages contains the page components for the application.
// This file defines the Home page and its sub-components.
package pages

import (
	"form_exer/web/shared" // Local package with shared components (Banner, Footer, Page)
	"github.com/rohanthewiz/element" // Third-party HTML builder library
)

// STRUCT DEFINITION with EMBEDDING
// Home represents the home page structure and data
type Home struct {
	// EMBEDDED STRUCT: shared.Page is embedded (no field name)
	// This is the MIXIN PATTERN - Home "inherits" all Page fields and methods
	// Home can now call h.Banner(), h.Footer(), and access h.Title
	shared.Page

	// Additional field specific to Home page
	// This demonstrates extending the base Page with page-specific data
	Heading string
}

// METHOD with VALUE RECEIVER and NAMED RETURN VALUE
// (h Home) - value receiver, method belongs to Home type
// (out string) - NAMED RETURN VALUE: the return variable is declared in the signature
//   This creates a variable 'out' that's automatically returned (though we don't use it here)
//   Named returns make code self-documenting and enable "naked returns"
func (h Home) Render() (out string) {
	// Create a new HTML builder instance
	// element.NewBuilder() returns a pointer to a Builder
	b := element.NewBuilder()

	// METHOD CHAINING: Build the HTML structure
	// b.Body() creates a <body> tag with inline CSS
	// .R() is a VARIADIC METHOD - accepts any number of arguments
	b.Body("style", "background-color:tan").R(
		// FUNCTION CALL: element.RenderComponents is a helper function
		// It takes a builder and multiple components, renders each component
		// This demonstrates the COMPOSITE PATTERN - combining multiple components
		element.RenderComponents(b,
			// METHOD CALL on EMBEDDED FIELD: h.Banner() works because Page is embedded
			// This is equivalent to h.Page.Banner(), but Go allows the shorthand
			h.Banner(), // Returns Banner struct from the embedded Page

			// STRUCT LITERAL: Creating a CatAdoptionHero instance inline
			// Since CatAdoptionHero is empty, we use {}
			CatAdoptionHero{},

			// Another method from the embedded Page
			h.Footer(), // Returns Footer struct
		),
		// Add a heading after the components
		// h.Heading accesses the Home struct's Heading field
		b.H1("style", "color:maroon;background-color:#dfc673").T(h.Heading),
	)

	// METHOD CALL: b.String() converts the builder to an HTML string
	// This returns the complete HTML document as a string
	return b.String()
}

// EMPTY STRUCT as a Component
// CatAdoptionHero has no fields but provides rendering behavior
// This is a stateless component - all its data is hardcoded in the Render method
type CatAdoptionHero struct{}

// METHOD with NAMED RETURN VALUE and 'any' TYPE
// (c CatAdoptionHero) - value receiver for the empty struct
// (dontCare any) - NAMED RETURN with type 'any' (alias for interface{})
//   The name "dontCare" documents that we ignore the return value
//   'any' can hold any type - maximum flexibility
func (c CatAdoptionHero) Render(b *element.Builder) (dontCare any) {
	// CONTAINER DIV with responsive design
	// max-width limits content width on large screens
	// margin:0 auto centers the container horizontally
	b.Div("style", "max-width:1200px; margin:0 auto; padding:40px 20px").R(
		// H2 heading - centered with custom styling
		b.H2("style", "text-align:center; color:#2c3e50; font-size:2.5em; margin-bottom:20px").T("Find Your Purr-fect Companion"),

		// P paragraph - .T() adds text content
		// Text can be a single string or multiple strings concatenated
		b.P("style", "text-align:center; color:#555; font-size:1.2em; margin-bottom:40px").T(
			"Give a loving cat a forever home. Browse our adoptable cats and kittens waiting to meet you!",
		),

		// CSS GRID LAYOUT: Modern, responsive card layout
		// display:grid creates a grid container
		// grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)) creates responsive columns:
		//   - auto-fit: automatically fits as many columns as possible
		//   - minmax(300px, 1fr): each column is min 300px, max 1 fraction of available space
		// gap:30px adds space between grid items
		b.Div("style", "display:grid; grid-template-columns:repeat(auto-fit, minmax(300px, 1fr)); gap:30px; margin-top:40px").R(
			// CARD COMPONENT PATTERN: Each cat is a card (Div) with image, text, and button
			// These 3 cards demonstrate repeating patterns - in real apps, use a loop with data

			// Cat Card 1 - Demonstrating the card structure
			b.Div("style", "background:white; border-radius:10px; box-shadow:0 4px 6px rgba(0,0,0,0.1); padding:20px").R(
				// IMG TAG with multiple attributes
				// Attributes are pairs: "name", "value", "name", "value"
				// This is a VARIADIC FUNCTION pattern - accepts any number of string pairs
				b.Img("src", "https://placekitten.com/400/300", "alt", "Orange tabby cat", "style", "width:100%; border-radius:8px; margin-bottom:15px"),

				// H3 heading for the cat's name
				b.H3("style", "color:#2c3e50; margin:10px 0").T("Whiskers"),

				// P paragraph with description
				// line-height:1.6 improves readability with proper spacing
				b.P("style", "color:#666; line-height:1.6").T("A friendly orange tabby who loves to play and cuddle. Great with kids and other pets. Age: 2 years."),

				// BUTTON element - demonstrates form controls
				// cursor:pointer changes cursor on hover (UX improvement)
				b.Button("style", "background-color:#e67e22; color:white; border:none; padding:10px 20px; border-radius:5px; cursor:pointer; font-size:1em; margin-top:10px").T("Meet Whiskers"),
			),

			// Cat 2
			b.Div("style", "background:white; border-radius:10px; box-shadow:0 4px 6px rgba(0,0,0,0.1); padding:20px").R(
				b.Img("src", "https://placekitten.com/401/300", "alt", "Gray and white cat", "style", "width:100%; border-radius:8px; margin-bottom:15px"),
				b.H3("style", "color:#2c3e50; margin:10px 0").T("Luna"),
				b.P("style", "color:#666; line-height:1.6").T("A calm and gentle gray beauty who enjoys quiet afternoons. Perfect for apartment living. Age: 4 years."),
				b.Button("style", "background-color:#e67e22; color:white; border:none; padding:10px 20px; border-radius:5px; cursor:pointer; font-size:1em; margin-top:10px").T("Meet Luna"),
			),

			// Cat 3
			b.Div("style", "background:white; border-radius:10px; box-shadow:0 4px 6px rgba(0,0,0,0.1); padding:20px").R(
				b.Img("src", "https://placekitten.com/402/300", "alt", "Black cat", "style", "width:100%; border-radius:8px; margin-bottom:15px"),
				b.H3("style", "color:#2c3e50; margin:10px 0").T("Shadow"),
				b.P("style", "color:#666; line-height:1.6").T("A playful black kitten full of energy and curiosity. Loves interactive toys and exploring. Age: 8 months."),
				b.Button("style", "background-color:#e67e22; color:white; border:none; padding:10px 20px; border-radius:5px; cursor:pointer; font-size:1em; margin-top:10px").T("Meet Shadow"),
			),
		),
	)

	// NAKED RETURN: Just "return" without a value
	// This works because we declared a named return value (dontCare any)
	// Go automatically returns the zero value of 'any', which is nil
	// Named returns enable this pattern, but use it sparingly for clarity
	return
}

// KEY CONCEPTS demonstrated in this file:
// 1. STRUCT EMBEDDING - Home embeds shared.Page for the mixin pattern
// 2. NAMED RETURN VALUES - Enables naked returns and self-documenting code
// 3. VALUE RECEIVERS - Methods receive copies of structs
// 4. METHOD CHAINING - Fluent API for building HTML
// 5. VARIADIC FUNCTIONS - Functions accepting any number of arguments
// 6. COMPOSITE PATTERN - Combining multiple components into pages
// 7. CSS GRID - Modern responsive layout directly in Go code
// 8. EMPTY STRUCTS - Zero-size structs for stateless components
// 9. 'any' TYPE - Go's universal type (interface{}) for maximum flexibility
// 10. NAKED RETURNS - Returning named values without explicit specification

