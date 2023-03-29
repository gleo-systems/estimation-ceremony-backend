// Package websocket provides functions to run a server handling WebSockets connection
package websocket

import (
	server "gleos/estimation/internal/app/gleos-estimation/websocket/internal"
)

func NewConfig(domain string, port uint) server.Config {
	return server.Config{Network: server.NetConfig{Domain: domain, Port: port}}
}

func RunServer(cfg server.Config) error {
	return server.Server{Cfg: cfg}.HandleConnections()
}
