package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
    application := app.Gerar()
   if err := application.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}