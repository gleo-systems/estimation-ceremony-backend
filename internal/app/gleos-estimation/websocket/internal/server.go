// Package server provides handlers for WebSockets incoming connections
package server

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"net/http"
	"strconv"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Config struct {
	Network NetConfig
}

type NetConfig struct {
	Domain string
	Port   uint
}

func (cfg NetConfig) getAddr() string {
	return cfg.Domain + ":" + uitoa(cfg.Port)
}

type Server struct {
	Cfg Config
}

func (srv Server) HandleConnections() (err error) {
	netAddr := srv.Cfg.Network.getAddr()
	log.Infow("Running WebSockets server", "address", netAddr)

	err = http.ListenAndServe(netAddr, http.HandlerFunc(EchoResponseHandler))
	if err != nil {
		log.Errorw("WebSockets server initialization failed", "err", err)
		return err
	}

	return nil
}

func EchoResponseHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Errorw("Cannot upgrade websocket connection", "err", err)
	}
	log.Debugw("Websocket connection succesfully upgraded")

	go func() {
		defer conn.Close()
		for {
			payload, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Errorw("Cannot read client data", "err", err)
				break
			}
			err = wsutil.WriteServerMessage(conn, op, payload)
			if err != nil {
				log.Errorw("Cannot write server data", "err", err)
				break
			}
		}
	}()
}

func uitoa(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}
