package server

import (
	"log"
	"net/http"
	"strings"
)

const (
	NETWORK_DOMAIN = "localhost"
	NETWORK_PORT   = "8000"
)

func initRouting() {
	http.HandleFunc("/", GetHomePage)
	log.Println("Server router successfully initilized")
}

func Run() {
	log.Println("Running application server...")
	initRouting()
	networkAddr := strings.Join([]string{NETWORK_DOMAIN, NETWORK_PORT}, "/")
	log.Fatal(
		http.ListenAndServe(networkAddr, nil))
}
