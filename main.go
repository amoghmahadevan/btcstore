package main

import (
	"log"
	"net/http"

	"github.com/amoghmahadevan/btcstore/handlers"
)

const LISTEN_ON = ":3000"

func main() {
	r := handlers.Router()

	log.Printf("Listening on %s", LISTEN_ON)
	log.Fatal(http.ListenAndServe(LISTEN_ON, r))
}
