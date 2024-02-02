package main

import (
	"log"
)

func main() {

	svc := NewCatFactService("https://catfact.ninja/fact")

	svc = NewLoggingService(svc)

	server := NewApiServer(":8000", svc)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
