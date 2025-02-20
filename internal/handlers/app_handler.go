package handlers

import "github.com/maxence-charriere/go-app/v10/pkg/app"

var GoAppHandler = &app.Handler{
	Name:        "Illuminate",
	Title:       "Illuminate",
	Description: "Illuminate browser extnsion home page",
	Icon:        app.Icon{Default: "/web/images/favicon.png", SVG: "/web/images/favicon.png"},

	Styles: []string{"/web/css/bootstrap.min.css", "/web/css/common.css", "/web/font/bootstrap-icons.min.css"},
	Scripts: []string{"/web/scripts/bootstrap.bundle.min.js",
		"/web/scripts/cash.min.js",
		"/web/scripts/common.js",
		"/web/scripts/commonmark.js",
	},
	Fonts: []string{"/web/font/fonts/bootstrap-icons.woff2"},
}
