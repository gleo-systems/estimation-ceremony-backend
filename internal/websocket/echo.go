package websocket

import (
	"estimation/ceremony/internal/log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

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
