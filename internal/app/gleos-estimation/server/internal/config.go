// Application server configuration
package server

import "strconv"

type Config struct {
	Network NetConfig
}

type NetConfig struct {
	Domain string
	Port   uint
}

func (cfg NetConfig) GetAddr() string {
	return cfg.Domain + ":" + uitoa(cfg.Port)
}

func uitoa(ui uint) string {
	return strconv.FormatUint(uint64(ui), 10)
}
