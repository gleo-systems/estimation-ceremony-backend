// Application command line runner
package main

import (
	"gleos/estimation/internal/app/server"
	"gleos/estimation/internal/pkg/log"

	goflags "github.com/jessevdk/go-flags"
)

type Options struct {
	Domain string `long:"domain" description:"Local network domain" required:"true"`
	Port   uint   `long:"port" description:"Local network port" required:"true"`
}

func main() {
	var opts Options
	if _, err := goflags.Parse(&opts); err != nil {
		panic(err)
	}
	config := server.NewConfig(opts.Domain, opts.Port)

	log.Infow("Starting application...", "params", opts)
	server.Run(config)
}
