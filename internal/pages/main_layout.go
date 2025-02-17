package pages

import (
	"github.com/ajaxe/illuminate-support2/internal/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type MainLayout struct {
	app.Compo
	Content app.UI
}

func (m *MainLayout) Render() app.UI {
	return m.container()
}

func (m *MainLayout) container() app.UI {
	return app.Div().Class("container-fluid").
		ID("main-container").
		Body(
			app.Div().Class("row").Body(
				app.Div().Class("col-lg-3 col-xl-2 bg-primary d-none d-lg-block").Body(
					m.leftPanel(),
				),
				app.Div().Class("col content-panel overflow-hidden").Body(
					m.rightPanel(),
				),
			),
		)
}
func (m *MainLayout) leftPanel() app.UI {
	app.Log("MainLayout.leftPanel")
	return components.AppNavLeftPanel()
}

func (m *MainLayout) rightPanel() app.UI {
	app.Log("MainLayout.rightPanel")
	return m.Content
}
