package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":7654", nil)
}
