package shared

import "github.com/rohanthewiz/element"

// Banner is a common header component for pages
type Banner struct {
	Title string
}

// Render implements element.Component interface
func (b Banner) Render(builder *element.Builder) any {
	builder.Header("style", "background-color:#2c3e50; color:white; padding:20px").R(
		builder.H1().T(b.Title),
	)
	return nil
}
