package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ajaxe/illuminate-support2/internal/handlers"
	"github.com/ajaxe/illuminate-support2/internal/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

var goAppHandler = &app.Handler{
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

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {

	app.Route("/", func() app.Composer { return &pages.HomePage{} })
	app.Route("/home", func() app.Composer { return &pages.HomePage{} })
	app.Route("/releases", func() app.Composer { return &pages.ReleasesPage{} })

	app.RunWhenOnBrowser()

	http.HandleFunc("/healthcheck", withLogging(&handlers.HealthCheckHandler{}))
	http.HandleFunc("/", withLogging(goAppHandler))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func withLogging(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer timer(r)()
		h.ServeHTTP(w, r)
	}
}
func timer(r *http.Request) func() {
	start := time.Now()
	return func() {
		log.Printf("addr=%s method=%s url=%s dur=%s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	}
}
