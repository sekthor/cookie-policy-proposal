package main

import (
	"log"

	"github.com/sekthor/cookie-policy-proposal/internal"
)

func main() {
	server := internal.Server{}

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
