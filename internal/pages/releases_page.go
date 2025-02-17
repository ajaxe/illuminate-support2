package pages

import (
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type ReleasesPage struct {
	app.Compo
}

func (h *ReleasesPage) OnNav(ctx app.Context) {
	ctx.Async(func() {
		res, err := http.Get("/web/content/releases.md")
		if err != nil {
			app.Log(err)
			return
		}
		defer res.Body.Close()
		if err != nil {
			b, _ := io.ReadAll(res.Body)
			app.Log(string(b))

			ctx.Dispatch(func(ctx app.Context) {
				app.Window().Get("showReleaseList").Invoke(string(b))
			})
		}
	})
}
func (h *ReleasesPage) Render() app.UI {
	return &MainLayout{
		Content: h.content(),
	}
}
func (h *ReleasesPage) content() app.UI {
	return app.Div().ID("releases").Class("section overflow-auto").
		Body(
			app.Div().Class("ms-auto me-auto content"),
		)

}
