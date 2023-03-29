// Application command line runner
package main

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"gleos/estimation/internal/app/gleos-estimation/websockets"

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
	config := websockets.NewConfig(opts.Domain, opts.Port)
	log.Infow("Application started by command", "options", opts)

	if err := websockets.RunServer(config); err != nil {
		panic(err)
	}
}
