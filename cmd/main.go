package main

import (
	"cmd/app"
	"log"
	"os"
)

func main() {
	app := app.Gerar()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
