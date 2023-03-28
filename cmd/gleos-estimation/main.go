// Application starting point
package main

import (
	"flag"
	"gleos/estimation/internal/app/server"
	"gleos/estimation/internal/pkg/log"
)

var netDomain = flag.String("domain", "localhost", "local network domain")
var netPort = flag.Int("port", 8000, "local network port")

func main() {
	config := server.NewConfig(*netDomain, *netPort)
	log.Infow("Starting gleos-estimation application", "config", config)
	server.Run(config)
}
