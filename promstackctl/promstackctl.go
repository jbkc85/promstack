package main

import (
	"log"

	"github.com/jbkc85/promstackctl/commands"
)

func main() {
	if err := commands.PromStackCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
