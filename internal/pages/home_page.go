package pages

import (
	"github.com/ajaxe/illuminate-support2/internal/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type HomePage struct {
	app.Compo
}

func (h *HomePage) Render() app.UI {
	return &MainLayout{
		Content: h.content(),
	}
}

func (h *HomePage) content() app.UI {
	return app.Div().ID("home").Class("section overflow-auto").
		Body(
			app.Div().Class("ms-auto me-auto content").
				Body(
					components.AppBanner(),
				),
			app.Div().Class("ms-auto me-auto content").
				Body(
					app.P().Class("lead").
						Text(`
							Apart from providing a way of highlighting and bookmarking
							phrases on the web, following are some of the features that make
							this extension more usable.
						`),
					app.P().Class("lead ms-3").
						ID("feature-list").
						Body(h.featureList()...),
					app.P().Class("pt-5").
						Body(
							app.A().Href("https://chrome.google.com/webstore/detail/illuminate/fehdmohlaoagebcpgpmhjfbhglnfednk").
								Target("_blank").
								Body(
									app.Img().Src("./web/images/iNEddTyWiMfLSwFD6qGq.png"),
								),
						),
				),
		)
}

func (h *HomePage) featureList() []app.UI {
	f := []string{
		"Customize color list for text highlighting.",
		"Disable extension on a single page or entire domain.",
		"Clear all text highlights with single click.",
		"See the summary of selections on the 'Options' page.",
	}
	elems := make([]app.UI, 0, (len(f)+1)*3)

	for _, v := range f {
		elems = append(elems, app.Raw(`<i class="bi bi-check2 text-success"></i>`))
		elems = append(elems, app.Span().Text(v))
		elems = append(elems, app.Br())
	}

	return elems
}
