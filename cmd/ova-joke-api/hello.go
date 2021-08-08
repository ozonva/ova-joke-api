package main

import (
	"log"
	"os"

	"github.com/ozonva/ova-joke-api/internal/app/hellower"
)

const (
	serviceName = "ova-joke-api"
)

func main() {
	if err := hellower.SayHelloFrom(os.Stdout, serviceName); err != nil {
		log.Fatalln(err)
	}
}
