// Application starting point
package main

import (
	"gleos/estimation/internal/app/server"
	"log"
)

func main() {
	log.Println("Starting gleos-estimation application...")
	server.Run()
}
