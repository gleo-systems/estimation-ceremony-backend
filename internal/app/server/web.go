// Application web server definition
package server

import (
	"fmt"
	"gleos/estimation/internal/pkg/log"
	"net/http"
)

func Run(cfg Config) {
	netAddr := cfg.getNetAddr()
	log.Infow("Running web server...", "network", netAddr)
	initRouting()
	log.Errorw("Web server initialization failed", "err", http.ListenAndServe(netAddr, nil))
}

func initRouting() {
	http.HandleFunc("/status", getStatus)
	log.Info("Web server router successfully initilized")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Web server is running")
}
