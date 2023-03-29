// Application server runner
package websockets

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"net/http"
	"strconv"
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
	log.Infow("Running WebSockets server...", "address", netAddr)

	if err = http.ListenAndServe(netAddr, nil); err != nil {
		log.Errorw("WebSockets server initialization failed", "err", err)
		return err
	}

	return nil
}

func uitoa(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}
