package websockets

import (
	websockets "gleos/estimation/internal/app/gleos-estimation/websockets/internal"
)

func NewConfig(domain string, port uint) websockets.Config {
	return websockets.Config{Network: websockets.NetConfig{Domain: domain, Port: port}}
}

func RunServer(cfg websockets.Config) error {
	return websockets.Server{Cfg: cfg}.HandleConnections()
}
