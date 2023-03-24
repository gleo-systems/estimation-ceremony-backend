package server

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "G-Leo Systems Estimation")
}
