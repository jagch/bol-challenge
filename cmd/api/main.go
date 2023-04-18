package main

import (
	"jagch/boletia/freecurrency/cmd/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
