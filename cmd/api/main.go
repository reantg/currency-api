package main

import (
	"log"
)

func main() {
	if err := runApp(); err != nil {
		log.Fatal("app run err", err)
	}
}
