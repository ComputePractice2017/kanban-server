package api

import (
	"fmt"
	"net/http"
)

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world,Beach")
}
