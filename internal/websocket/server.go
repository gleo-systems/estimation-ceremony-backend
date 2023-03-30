// Package websocket provides WebSockets server implementation to handle clients connections
package websocket

import (
	"estimation-ceremony/internal/log"
	"net/http"
	"strconv"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func NewConfig(netDomain string, netPort uint) Config {
	return Config{netDomain, netPort}
}

type Config struct {
	netDomain string
	netPort   uint
}

func (cfg Config) GetNetAddr() string {
	return cfg.netDomain + ":" + uitoa(cfg.netPort)
}

func NewServer(cfg Config) Server {
	return Server{cfg: cfg}
}

type Server struct {
	cfg Config
}

func (srv Server) HandleConnections() (err error) {
	netAddr := srv.cfg.GetNetAddr()
	log.Infow("Running WebSockets server", "address", netAddr)

	err = http.ListenAndServe(netAddr, http.HandlerFunc(echoResponseHandler))
	if err != nil {
		log.Errorw("WebSockets server initialization failed", "err", err)
		return err
	}

	return nil
}

func echoResponseHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Errorw("Cannot upgrade to websocket connection", "err", err)
	}
	log.Debug("Established new client websocket connection")

	go func() {
		defer conn.Close()
		for {
			payload, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Errorw("Cannot read client data", "err", err)
				break
			}
			log.Debugw("Read client data", "payload", string(payload))

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
