// Application basic web handlers
package server

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GLEOS Estimation Home Page")
}
