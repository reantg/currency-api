package main

import (
	"log"
)

// TODO rename app -> api

func main() {
	if err := runApp(); err != nil {
		log.Fatal("app run err", err)
	}
}
