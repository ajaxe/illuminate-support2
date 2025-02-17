package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type NavLeftPanel struct {
	app.Compo
}

func (n *NavLeftPanel) Render() app.UI {
	return app.Div().Class("d-flex flex-column justify-content-center align-items-center side-panel text-white").
		Body(
			AppLogo(),
			n.appName(),
			AppNavBarItems(NavListOptions{TextColor: "text-white"}),
		)

}
func (n *NavLeftPanel) appName() app.UI {
	return app.Div().Class("text-center").
		Body(app.H1().Text("Illuminate"))
}
