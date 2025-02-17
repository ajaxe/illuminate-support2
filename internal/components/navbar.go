package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type NavBar struct {
	app.Compo
}

func (h *NavBar) Render() app.UI {
	return app.Nav().
		Class("fixed-top navbar navbar-expand-lg navbar-dark bg-primary d-lg-none").
		Body(h.logo(), h.brandName(), h.navToggler(), h.navItems())
}

func (n *NavBar) logo() app.UI {
	return AppLogo()
}
func (n NavBar) brandName() app.UI {
	return app.A().
		Class("navbar-brand").
		Href("#").
		Text("Illuminate")
}
func (n NavBar) navItems() app.UI {
	return app.Div().
		Class("collapse navbar-collapse").
		ID("navbarSupportedContent").
		Body(AppNavBarItems(NavListOptions{
			TextColor: "text-white",
			ListCSS: "nav flex-column font-weight-bold",
		}))
}
func (n NavBar) navToggler() app.UI {
	return app.Button().
		Class("navbar-toggler").
		Class("collapsed").
		Type("button").
		DataSet("bs-toggle", "collapse").
		DataSet("bs-target", "#navbarSupportedContent").
		Aria("controls", "navbarSupportedContent").
		Aria("expanded", "false").
		Aria("label", "Toggle navigation").
		Body(
			app.Span().Class("navbar-toggler-icon"),
		)
}
