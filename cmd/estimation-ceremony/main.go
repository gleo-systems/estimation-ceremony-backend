// The estimation-ceremony command implementation running application server
package main

import (
	"gleos/estimation/internal/app/gleos-estimation/log"
	"gleos/estimation/internal/app/gleos-estimation/websocket"

	// akokot - consider github.com/spf13/cobra replacement
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

	srv := websocket.NewServer(config)
	if err := srv.HandleConnections(); err != nil {
		panic(err)
	}
}
