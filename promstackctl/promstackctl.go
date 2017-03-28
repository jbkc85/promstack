package main

import (
	"log"

	"scm01cou-dv.col.sp/promstackctl/commands"
)

func main() {
	if err := commands.PromStackCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
