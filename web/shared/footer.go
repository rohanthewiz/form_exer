package shared

import "github.com/rohanthewiz/element"

// Footer is a common footer component for pages
type Footer struct{}

// Render implements element.Component interface
func (f Footer) Render(b *element.Builder) any {
	b.Div("style", "background-color:lightgray").R(
		b.P("style", "color:gray").T("Copyright &copy; 2025"),
	)
	return nil
}
