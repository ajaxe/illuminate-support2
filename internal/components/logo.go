package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Logo struct {
	app.Compo
}

func (l *Logo) Render() app.UI {
	return app.Div().Class("logo")
}
