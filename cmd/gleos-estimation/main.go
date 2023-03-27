// Application starting point
package main

import (
	"gleos/estimation/internal/app/server"
	"gleos/estimation/internal/pkg/log"
)

func main() {
	log.Info("Starting gleos-estimation application")
	server.Run()
}
