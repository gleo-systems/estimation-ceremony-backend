// Application web server configuration
package server

import "strconv"

type Config struct {
	netDomain string
	netPort   int
}

func (cfg Config) getNetAddr() string {
	return cfg.netDomain + ":" + strconv.Itoa(cfg.netPort)
}

func NewConfig(netDomain string, netPort int) Config {
	return Config{netDomain, netPort}
}
