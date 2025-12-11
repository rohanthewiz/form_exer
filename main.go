// Package main is a special package in Go that defines an executable program.
// When you build a "package main", Go creates an executable binary file.
// This package demonstrates the usage of rweb, a lightweight HTTP web server framework for Go.
// It showcases various features including routing, middleware, static files, file uploads,
// Server-Sent Events (SSE), and reverse proxy capabilities.
package main

// Import declarations bring other packages into this file's scope.
// Go organizes imports into groups (standard library, then third-party packages).
import (
	// Standard library imports (built into Go)
	"fmt"    // Package for formatted I/O (printing, string formatting)
	"io"     // Package for I/O primitives (reading, writing)
	"log"    // Package for simple logging
	"net/http" // Package for HTTP client and server implementations
	"os"     // Package for operating system functionality (file operations)
	"strings" // Package for string manipulation

	// Local package imports (from this module)
	"form_exer/web/pages" // Our page components (HomePage, Contact, etc.)

	// Third-party package imports (external dependencies defined in go.mod)
	"github.com/rohanthewiz/element" // HTML element builder library
	"github.com/rohanthewiz/rweb"    // Lightweight web framework
)

// main() is the entry point of the program. Go automatically calls this function when the program starts.
// Every executable Go program must have exactly one main() function in the main package.
func main() {
	// DEFER: The defer keyword schedules a function call to run after the surrounding function returns.
	// Deferred functions execute in LIFO (Last In, First Out) order.
	// Note: defer does not fire on CTRL-C because CTRL-C uses os.Exit() for immediate shutdown
	defer func() {
		fmt.Println("Exiting main()...")
	}() // The () at the end immediately invokes this anonymous function (but defer delays its execution)

	// STRUCT LITERAL with NAMED FIELDS: Creating a new rweb server instance
	// rweb.ServerOptions is a struct type, and we're creating an instance using a struct literal.
	// Named fields (Address: value) make the code self-documenting and allow fields in any order.
	s := rweb.NewServer(rweb.ServerOptions{
		// Address specifies the TCP address for the server to listen on
		// Format: ":port" listens on all network interfaces (0.0.0.0:8000)
		// This is preferred over "localhost:8000" for Docker compatibility
		Address: ":8000",

		// Verbose is a boolean field that enables detailed request/response logging
		// Go's zero value for bool is false, so we explicitly set it to true
		Verbose: true,

		// Debug is another boolean field for additional debugging information
		// Explicitly setting to false for clarity (same as omitting it)
		Debug: false,
	})

	// SHORT VARIABLE DECLARATION: The := operator declares and initializes a variable
	// Go infers the type from the right-hand side (here: *rweb.Server)
	// This is equivalent to: var s *rweb.Server = rweb.NewServer(...)

	// METHOD CALL: Calling the Use() method on the server instance
	// Use() registers middleware that runs before route handlers
	// rweb.RequestInfo is a pre-built middleware function provided by the rweb package
	s.Use(rweb.RequestInfo)

	/*	// MIDDLEWARE PATTERN: Middleware are functions that process requests before they reach handlers
		// Middleware 1: Request logging middleware
		// This middleware logs each request's method, path, response status, and duration
		// Middleware in rweb is executed in the order it's registered
		s.Use(func(ctx rweb.Context) error {
			// Record the start time of the request
			start := time.Now()

			// defer ensures this runs after the request is handled
			// This allows us to capture the final response status and calculate duration
			defer func() {
				// Log format: GET "/path" -> 200 [150ms]
				fmt.Printf("In Midware 1 - %s %q -> %d [%s]\n",
					ctx.Request().Method(),  // HTTP method (GET, POST, etc.)
					ctx.Request().Path(),    // Request path
					ctx.Response().Status(), // Response status code
					time.Since(start))       // Request duration
			}()

			// Call ctx.Next() to pass control to the next middleware or handler
			// This is crucial - without it, the request chain stops here
			return ctx.Next()
		})
	*/

	// type MidWare func(ctx rweb.Context) error
	// var authMidWare rweb.Handler

	_ = func(ctx rweb.Context) error {
		fmt.Println("**-> Checking Auth...")

		reqPath := ctx.Request().Path()
		if strings.Contains(reqPath, "roh") {
			fmt.Println("**-> Auth OK")
			return ctx.Next()
		}

		ctx.Response().SetStatus(http.StatusUnauthorized) // 401
		return nil
	}

	// s.Use(authMidWare)

	/*	// We could put the middleware function definition in a variable like this
		midWare2 := func(ctx rweb.Context) error {
			fmt.Println("In MidWare 2: ", ctx.Request().Method(), ctx.Request().Path())
			// Always call ctx.Next() unless you want to stop the request chain
			return ctx.Next()
		}

		// Middleware 2: Simple demonstration middleware
		// Shows that multiple middleware can be chained together
		s.Use(midWare2)
	*/

	/*	s.Use(func(ctx rweb.Context) error {
			fmt.Println("In MidWare 3: ", ctx.Request().Host())
			return ctx.Next()
		})
	*/

	// ===== HTTP ROUTE HANDLERS =====
	// Routes map URL paths to handler functions
	// The server's router (master handler) looks up which route matches the incoming request

	// ROUTE HANDLER with ANONYMOUS FUNCTION
	// s.Get() registers a handler for GET requests to the "/" path
	// The second argument is an anonymous function (also called a function literal or lambda)
	// FUNCTION SIGNATURE: func(ctx rweb.Context) error
	//   - Takes one parameter: ctx (context with request/response data)
	//   - Returns an error (nil means success, non-nil error is handled by the framework)
	s.Get("/", func(ctx rweb.Context) error {
		// METHOD CHAINING: ctx.Response() returns a response object, then we call SetHeader() on it
		// SetHeader sets an HTTP response header (key-value pair)
		ctx.Response().SetHeader("Content-Type", "text/html; charset=utf-8")

		// CALLING METHODS ACROSS PACKAGES
		// pages.HomePage is a struct instance from the pages package
		// We call its Render() method, which returns an HTML string
		// ctx.WriteHTML() sends that HTML back to the client
		// The return statement returns the error (or nil) from WriteHTML
		return ctx.WriteHTML(pages.HomePage.Render())
	})

	// Another GET route - same pattern as above
	// This demonstrates that we can have multiple routes with different paths
	// GET requests typically retrieve and display data (idempotent - safe to repeat)
	s.Get("/contact", func(ctx rweb.Context) error {
		ctx.Response().SetHeader("Content-Type", "text/html; charset=utf-8")
		// pages.Contact is another page instance, similar to HomePage
		return ctx.WriteHTML(pages.Contact.Render())
	})

	/*	s.Get("/roh", func(ctx rweb.Context) error {
			ctx.Response().SetHeader("Content-Type", "text/plain; charset=utf-8")

			// WriteString sends a plain text response
			return ctx.WriteString("Welcome to Roh!\n")
		})

		// Route parameters demonstration
		// The radix tree router correctly distinguishes between parameterized and static routes
		// Test with: curl http://localhost:8080/greet/John
		s.Get("/greet/:name", func(ctx rweb.Context) error {
			// Access route parameters using ctx.Request().Param("paramName")
			// The :name parameter captures any value in that URL segment
			return ctx.WriteString("Hello " + ctx.Request().PathParam("name"))
		})
	*/

	// POST ROUTE with ROUTE PARAMETERS
	// POST requests typically modify data on the server (non-idempotent - side effects)
	// ROUTE PARAMETERS: ":form_id" is a URL parameter that captures any value in that position
	// Example: /post-form-data/123 → form_id = "123"
	// Example: /post-form-data/abc → form_id = "abc"
	// Test with: curl -X POST http://localhost:8080/post-form-data/123 -d "dept=engineering&name=JohnDoe"
	s.Post("/post-form-data/:form_id",
		func(ctx rweb.Context) error {
			// FORM DATA EXTRACTION: FormValue() retrieves data from POST request body
			// These values come from the form data sent in the request
			dept := ctx.Request().FormValue("dept")      // form field "dept=engineering"
			formId := ctx.Request().PathParam("form_id") // URL path parameter "123"
			name := ctx.Request().FormValue("name")      // form field "name=JohnDoe"

			// STRING FORMATTING: fmt.Sprintf() works like printf - formats a string
			// %s is a placeholder for string values
			outStr := fmt.Sprintf("Posted - form_id: %s, dept: %s, name: %s", formId, dept, name)

			// WriteString sends plain text back to the client
			return ctx.WriteString(outStr)
		})

	// POST route for contact form submission
	// This handles the form data from the contact page
	s.Post("/contact",
		func(ctx rweb.Context) error {
			// Extract multiple form fields from the POST request
			name := ctx.Request().FormValue("name")       // form field "name"
			email := ctx.Request().FormValue("email")     // form field "email"
			message := ctx.Request().FormValue("message") // form field "message"
			outStr := fmt.Sprintf("Posted - name: %s, email: %s, message: %s", name, email, message)

			// FLUENT API / METHOD CHAINING: Building HTML dynamically
			// element.NewBuilder() creates a new HTML builder
			b := element.NewBuilder()

			// METHOD CHAINING with VARIADIC FUNCTIONS
			// Body() creates a <body> tag with style attribute
			// R() is a variadic function - it accepts any number of arguments (components)
			// Each method returns the builder, allowing us to chain calls
			b.Body("style", "background-color:darkgreen").R(
				// H1() creates an <h1> tag, T() adds text content
				b.H1("style", "color:maroon;background-color:#dfc673").T("Welcome"),
				b.Hr(), // Hr() creates an <hr> horizontal rule tag
				b.P().T(outStr), // P() creates a <p> paragraph tag
			)

			// String() converts the builder to an HTML string
			return ctx.WriteHTML(b.String())
		})

	// STATIC FILE SERVING
	// StaticFiles() serves files from the filesystem
	// Parameters: (URL prefix, filesystem path, segments to strip)
	// Example: Request to "/.well-known/some-file.txt" → serves file at "/some-file.txt"
	s.StaticFiles("/.well-known/", "/", 0)

	// FILE UPLOAD HANDLER
	// Demonstrates handling multipart/form-data (file uploads + regular form fields)
	// Test with: curl -X POST -F "vehicle=car" -F "file=@somefile.txt" http://localhost:8080/upload
	s.Post("/upload", func(c rweb.Context) error {
		// Get the request object for convenience
		req := c.Request()

		// MULTIPART FORM: Can contain both regular fields and file uploads
		// Extract regular form field (not a file)
		name := req.FormValue("vehicle")
		fmt.Println("vehicle:", name)

		// MULTIPLE RETURN VALUES: Go functions can return multiple values
		// GetFormFile returns 3 values: (file, fileHeader, error)
		// The BLANK IDENTIFIER (_) ignores the second return value (fileHeader)
		// This is Go's way of explicitly discarding values we don't need
		file, _, err := req.GetFormFile("file")

		// ERROR HANDLING PATTERN: Check if err is not nil
		// In Go, errors are values and must be explicitly checked
		// If there's an error, return it immediately (early return pattern)
		if err != nil {
			return err
		}

		// DEFER for RESOURCE CLEANUP: Ensure file is closed when function exits
		// This prevents resource leaks even if the function returns early or panics
		// defer runs in LIFO order (Last In, First Out)
		defer file.Close()

		// Read the entire file content into a byte slice ([]byte)
		// io.ReadAll reads until EOF (End Of File) or error
		// Note: For large files, consider streaming to disk instead of loading into memory
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// OCTAL LITERAL: 0666 is an octal number (base 8) representing file permissions
		// In Unix: 0666 = rw-rw-rw- (read/write for owner, group, others)
		// os.WriteFile creates (or overwrites) a file with the given data and permissions
		err = os.WriteFile("uploaded_file.txt", data, 0666)
		if err != nil {
			return err
		}

		// Returning nil indicates success (no error occurred)
		// In Go, functions that return error typically return nil on success
		return nil
	})

	// SERVER STARTUP
	// s.Run() starts the HTTP server and blocks until shutdown
	// It returns an error if the server fails to start or crashes
	// log.Println() prints the error (if any) when the server stops
	// This is the last line of main() - the program waits here while serving requests
	log.Println(s.Run())
}

// ===== EXAMPLE TEST OUTPUT =====
// The comments below show example curl commands and their outputs
// These demonstrate how the API endpoints work in practice

// Outputs
// >curl -X POST -d "dept=support" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123%
// >curl -X POST -d "dept=support" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123, dept: support%
// >curl -X POST -d "dept=support" -d "name=Sue" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123, dept: support%
// >curl -X POST -d "dept=support" -d "name=Sue" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/roh
// Posted - form_id: 123, dept: support, name:Sue%
