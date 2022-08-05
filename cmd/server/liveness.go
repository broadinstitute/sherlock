package main

import (
	"fmt"
	"net/http"
)

type livenessHandler struct{}

func (h livenessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "OK")
}
