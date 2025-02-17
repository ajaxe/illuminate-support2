package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Banner struct {
	app.Compo
}

func (b *Banner) Render() app.UI {
	return app.Div().Class("container-fluid p-5 mb-4 bg-body-tertiary rounded-3").
		Body(
			app.Div().Class("py-5").
				Body(
					app.H1().Class("display-3").Text("Illuminate"),
					app.Blockquote().Class("blockquote").Body(
						app.P().Class("mb-0").Text("Provide a way for an avid reader to bookmark phrases across the Internet."),
					),
				),
		)
}
