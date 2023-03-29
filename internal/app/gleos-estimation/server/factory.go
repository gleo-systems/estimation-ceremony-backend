package server

import (
	server "gleos/estimation/internal/app/gleos-estimation/server/internal"
)

type Server interface {
	Run()
}

func NewInstance(cfg server.Config) Server {
	return server.WebSocketsServer{Cfg: cfg}
}

func NewConfig(domain string, port uint) server.Config {
	return server.Config{Network: server.NetConfig{Domain: domain, Port: port}}
}
