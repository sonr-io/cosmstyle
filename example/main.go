package main

import (
	"net/http"
	"time"

	"github.com/sonr-io/cosmstyle/example/pages"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(pages.TimeComponent(time.Now())))
	http.Handle("/404", templ.Handler(pages.NotFoundComponent(), templ.WithStatus(http.StatusNotFound)))
	http.ListenAndServe(":8080", nil)
}
