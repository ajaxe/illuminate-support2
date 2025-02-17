package pages

import (
	"io"
	"net/http"

	"github.com/ajaxe/illuminate-support2/internal/helpers"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type ReleasesPage struct {
	app.Compo
}

func (h *ReleasesPage) OnNav(ctx app.Context) {
	baseURL := app.Window().URL()
	baseURL.Path = "/web/content/releases.md"

	u := baseURL.String()

	ctx.Async(func() {
		helpers.AppLogf("Fetching releases from %s", u)
		res, err := http.Get(u)
		if err != nil {
			app.Log(err)
			return
		}
		defer res.Body.Close()

		b, _ := io.ReadAll(res.Body)
		helpers.AppLog("fetch complete")

		ctx.Dispatch(func(ctx app.Context) {
			app.Window().Get("showReleaseList").Invoke(string(b))
		})
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
