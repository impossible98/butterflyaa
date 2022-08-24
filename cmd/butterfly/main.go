package main

import (
	// import built-in packages
	"fmt"
	"os"
	// import local packages
	"github.com/impossible98/butterfly/global"
	"github.com/impossible98/butterfly/internal/command"
)

func logo() {
	const line string = "----------------------------------------------------"
	fmt.Println(global.ASCII_NAME + "\n" + line + "\n\tv" + global.VERSION + "\n" + line)
}

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "serve" {
			logo()
			command.Serve()
		} else if os.Args[1] == "--version" {
			command.Version()
		}
	} else if len(os.Args) == 1 {
		command.Version()
		command.Help()
	}
}
