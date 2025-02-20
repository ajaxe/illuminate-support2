package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/ajaxe/illuminate-support2/internal/handlers"
)

func main() {
	// ...
	// Generate the static files
	// ...
	outDir := "../../dist/static"
	resources := []string{"/", "/app.css", "/app-worker.js", "/app.js", "/manifest.webmanifest", "/wasm_exec.js"}

	server := httptest.NewServer(handlers.GoAppHandler)
	defer server.Close()

	fmt.Printf("test server url: %s\n", server.URL)

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.MkdirAll(outDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	for _, r := range resources {
		fname := r
		if r == "/" {
			fname = "/index.html"
		}
		f, err := os.Create(outDir + fname)
		if err != nil {
			panic(err)
		}
		fetchResource(server.URL, r, f)
	}

}

func fetchResource(url, resource string, f io.Writer) {
	resp, err := http.Get(url + resource)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	c, err := io.Copy(f, resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Printf("resource '%-25s' Copied %d bytes\n", resource, c)
}
