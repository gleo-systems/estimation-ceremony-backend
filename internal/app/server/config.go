// Application server configuration
package server

import "strconv"

func NewConfig(domain string, port uint) config {
	return config{network: netConfig{domain, port}}
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
