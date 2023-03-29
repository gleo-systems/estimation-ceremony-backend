// Application server runner
package server

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"net/http"
)

type WebSocketsServer struct {
	Cfg Config
}

func (srv WebSocketsServer) Run() {
	netAddr := srv.Cfg.Network.GetAddr()
	log.Infow("Running WebSockets server...", "address", netAddr)

	err := http.ListenAndServe(netAddr, nil)
	if err != nil {
		log.Errorw("WebSockets server initialization failed", "err", err)
	}
}
