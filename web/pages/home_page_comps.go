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
			h.Banner(),         // Common banner from Page mixin
			CatAdoptionHero{}, // Cat adoption content
			h.Footer(),         // Common footer from Page mixin
		),
		b.H1("style", "color:maroon;background-color:#dfc673").T(h.Heading),
	)

	return b.String()
}

// CatAdoptionHero component renders the cat adoption content
type CatAdoptionHero struct{}

func (c CatAdoptionHero) Render(b *element.Builder) (dontCare any) {
	b.Div("style", "max-width:1200px; margin:0 auto; padding:40px 20px").R(
		b.H2("style", "text-align:center; color:#2c3e50; font-size:2.5em; margin-bottom:20px").T("Find Your Purr-fect Companion"),
		b.P("style", "text-align:center; color:#555; font-size:1.2em; margin-bottom:40px").T(
			"Give a loving cat a forever home. Browse our adoptable cats and kittens waiting to meet you!",
		),

		// Cat gallery
		b.Div("style", "display:grid; grid-template-columns:repeat(auto-fit, minmax(300px, 1fr)); gap:30px; margin-top:40px").R(
			// Cat 1
			b.Div("style", "background:white; border-radius:10px; box-shadow:0 4px 6px rgba(0,0,0,0.1); padding:20px").R(
				b.Img("src", "https://placekitten.com/400/300", "alt", "Orange tabby cat", "style", "width:100%; border-radius:8px; margin-bottom:15px"),
				b.H3("style", "color:#2c3e50; margin:10px 0").T("Whiskers"),
				b.P("style", "color:#666; line-height:1.6").T("A friendly orange tabby who loves to play and cuddle. Great with kids and other pets. Age: 2 years."),
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

	return
}

