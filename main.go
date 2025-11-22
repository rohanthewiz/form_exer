// Package main demonstrates the usage of rweb, a lightweight HTTP web server framework for Go.
// This example showcases various features including routing, middleware, static files, file uploads,
// Server-Sent Events (SSE), and reverse proxy capabilities.
package main

import (
	"fmt"
	"form_exer/web/pages"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func main() {
	// Defers does not fire on CTRL-C bc CTRL-C uses os.Exit() // immediate shutdown
	defer func() {
		fmt.Println("Bye")
	}()

	// Create a new rweb server instance with configuration options
	s := rweb.NewServer(rweb.ServerOptions{
		// Listen on port 8080 on all network interfaces
		// Use ":8080" format (not "localhost:8080") for Docker compatibility
		Address: ":8000",

		// Enable verbose logging to see detailed request/response information
		Verbose: true,

		// Debug mode is disabled (would show additional debugging information)
		Debug: false,
	})

	s.Use(rweb.RequestInfo)
	/*	// Middleware 1: Request logging middleware
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

	// We could put the middleware function definition in a variable like this
	midWare2 := func(ctx rweb.Context) error {
		fmt.Println("In MidWare 2: ", ctx.Request().Method(), ctx.Request().Path())
		// Always call ctx.Next() unless you want to stop the request chain
		return ctx.Next()
	}

	// Middleware 2: Simple demonstration middleware
	// Shows that multiple middleware can be chained together
	s.Use(midWare2)

	s.Use(func(ctx rweb.Context) error {
		fmt.Println("In MidWare 3: ", ctx.Request().Host())
		return ctx.Next()
	})

	// [master handler -> Lookup route]

	// Route: Root endpoint
	// Handles GET requests to "/" - the home page
	// Test with: curl http://localhost:8080/
	s.Get("/", func(ctx rweb.Context) error {
		ctx.Response().SetHeader("Content-Type", "text/html; charset=utf-8")
		return ctx.WriteHTML(pages.HomePage.Render())
	})

	s.Get("/roh", func(ctx rweb.Context) error {
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

	// POST request with route parameter
	// Demonstrates that route parameters work with all HTTP methods
	// Test with: curl -X POST http://localhost:8080/post-form-data/123 -d "dept=engineering&name=JohnDoe"
	s.Post("/post-form-data/:form_id",
		func(ctx rweb.Context) error {
			dept := ctx.Request().FormValue("dept")      // form data "dept=engineering"
			formId := ctx.Request().PathParam("form_id") // route/path parameter "123"
			name := ctx.Request().FormValue("name")      // form data "name=JohnDoe"
			outStr := fmt.Sprintf("Posted - form_id: %s, dept: %s, name: %s", formId, dept, name)

			return ctx.WriteString(outStr)
		})

	s.Post("/contact",
		func(ctx rweb.Context) error {
			name := ctx.Request().FormValue("name") // form data "name=JohnDoe"
			email := ctx.Request().FormValue("email")
			message := ctx.Request().FormValue("message")
			outStr := fmt.Sprintf("Posted - name: %s, email: %s, message: %s", name, email, message)

			b := element.NewBuilder()
			b.Body("style", "background-color:darkgreen").R(
				b.H1("style", "color:maroon;background-color:#dfc673").T("Welcome"),
				b.Hr(),
				b.P().T(outStr),
			)

			return ctx.WriteHTML(b.String())
		})

	// Example 3: Serve .well-known files (for SSL certificates, etc.)
	// URL: http://localhost:8080/.well-known/some-file.txt
	// Maps to: /some-file.txt (strips 0 segments, keeps full path)
	s.StaticFiles("/.well-known/", "/", 0)

	// File upload handler
	// Handles multipart/form-data POST requests
	// Test with: curl -X POST -F "vehicle=car" -F "file=@somefile.txt" http://localhost:8080/upload
	s.Post("/upload", func(c rweb.Context) error {
		req := c.Request()

		// Extract regular form fields from the multipart form
		name := req.FormValue("vehicle")
		fmt.Println("vehicle:", name)

		// Get the uploaded file
		// GetFormFile returns: file handle, file header (with metadata), error
		// We're ignoring the file header (second return value) here
		file, _, err := req.GetFormFile("file")
		if err != nil {
			return err
		}
		// Always close the file when done
		defer file.Close()

		// Read the entire file content into memory
		// For large files, consider streaming to disk instead
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// Save the uploaded file to disk
		// 0666 permissions: read/write for owner, group, and others
		err = os.WriteFile("uploaded_file.txt", data, 0666)
		if err != nil {
			return err
		}

		// Return nil indicates successful handling
		return nil
	})

	// Start the HTTP server
	// This blocks until the server is shut down
	// log.Fatal ensures any startup errors are logged before exiting
	log.Fatal(s.Run())
}

// Outputs
// >curl -X POST -d "dept=support" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123%
// >curl -X POST -d "dept=support" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123, dept: support%
// >curl -X POST -d "dept=support" -d "name=Sue" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/123
// Posted - form_id: 123, dept: support%
// >curl -X POST -d "dept=support" -d "name=Sue" -H "Content-Type: application/x-www-form-urlencoded" http://localhost:8000/post-form-data/roh
// Posted - form_id: 123, dept: support, name:Sue%
