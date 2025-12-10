// Package pages contains page component definitions and instances.
// Each page in the application is defined here with its data and rendering logic.
package pages

// Import the shared package to access the Page mixin
import "form_exer/web/shared"

// PACKAGE-LEVEL VARIABLE: HomePage is a global variable accessible throughout the program
// EXPORTED VARIABLE: Starts with capital letter, so it's accessible from other packages
// SINGLETON PATTERN: There's only one HomePage instance for the entire application
//
// STRUCT LITERAL with EMBEDDED FIELD
// Home is defined in home_page_comps.go - this creates an instance of it
var HomePage = Home{
	// EMBEDDED FIELD: Page is embedded (no field name, just the type)
	// This gives Home access to all Page fields and methods
	// We initialize it with a nested struct literal
	Page: shared.Page{Title: "My Website"},

	// Regular field: Heading is a specific field of the Home struct
	// This is different from Page.Title - Heading is used for page content
	Heading: "Home Page",
}

// KEY CONCEPTS demonstrated in this file:
// 1. PACKAGE-LEVEL VARIABLES - var at package level creates global variables
// 2. EXPORTED vs UNEXPORTED - HomePage (exported) can be used in other packages
// 3. SINGLETON PATTERN - Single instance created at program startup
// 4. STRUCT EMBEDDING - Page is embedded, giving Home all its functionality
// 5. NESTED STRUCT LITERALS - shared.Page{...} is nested inside Home{...}
// 6. INITIALIZATION ORDER - This runs before main(), during package initialization
