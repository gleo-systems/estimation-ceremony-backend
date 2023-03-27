// Application web server configuration
package server

import (
	"gleos/estimation/internal/pkg/log"
	"net/http"
)

const (
	_NETWORK_DOMAIN = "localhost"
	_NETWORK_PORT   = "8000"
)

func Run() {
	networkAddr := _NETWORK_DOMAIN + ":" + _NETWORK_PORT
	log.Infow("Running web server...", "network", networkAddr)
	initRouting()
	log.Errorw("Web server initialization failed", "err", http.ListenAndServe(networkAddr, nil))
}

func initRouting() {
	http.HandleFunc("/", GetHomePage)
	log.Info("Web server router successfully initilized")
}
