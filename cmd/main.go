package main

import (
	"hal9k/cmd/app"
	"log"
)

func main() {
	cmd := app.NewHal9000Command()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
