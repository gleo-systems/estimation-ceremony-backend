// Package websocket provides WebSockets server implementation to handle full-duplex communication
package websocket

import (
	"estimation/ceremony/internal/log"
	"net/http"
	"strconv"
)

func NewConfig(host string, port int) Config {
	return Config{host, port}
}

type Config struct {
	host string
	port int
}

func NewServer(cfg Config) Server {
	return Server{cfg: cfg}
}

type Server struct {
	cfg Config
}

func (srv Server) HandleConnections() (err error) {
	addr := srv.cfg.host + ":" + strconv.Itoa(srv.cfg.port)
	log.Infow("Running websockets server", "address", addr)

	router := initRouting()
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Errorw("Websockets server initialization failed", "err", err)
		return err
	}
	return
}

// TODO: akokot - consider using WEB framework github.com/labstack/echo/v4
func initRouting() (router *http.ServeMux) {
	router = http.NewServeMux()
	router.HandleFunc("/echo", echoResponseHandler)
	return
}
