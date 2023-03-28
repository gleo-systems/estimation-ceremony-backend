// Application server configuration
package server

import "strconv"

func NewConfig(netDomain string, netPort uint) config {
	return config{network: netConfig{netDomain, netPort}}
}

type config struct {
	network netConfig
}

type netConfig struct {
	domain string
	port   uint
}

func (cfg netConfig) getAddr() string {
	return cfg.domain + ":" + uitoa(cfg.port)
}

func uitoa(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}
