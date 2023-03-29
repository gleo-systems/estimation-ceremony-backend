// The gleos-estimation command runs the application
package main

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"gleos/estimation/internal/app/gleos-estimation/websocket"

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
	config := websocket.NewConfig(opts.Domain, opts.Port)
	log.Infow("Executed gleos-estimation command", "options", opts)

	if err := websocket.RunServer(config); err != nil {
		panic(err)
	}
}
