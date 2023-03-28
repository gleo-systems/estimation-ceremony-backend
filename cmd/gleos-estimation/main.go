// Application command line runner
package main

import (
	"gleos/estimation/internal/app/server"
	"gleos/estimation/internal/pkg/log"

	goflags "github.com/jessevdk/go-flags"
)

type Parameters struct {
	Domain string `long:"domain" description:"Local network domain" required:"true"`
	Port   uint   `long:"port" description:"Local network port" required:"true"`
}

func main() {
	var params Parameters
	if _, err := goflags.Parse(&params); err != nil {
		panic(err)
	}
	config := server.NewConfig(params.Domain, params.Port)

	log.Infow("Starting application", "params", params)
	server.Run(config)
}
