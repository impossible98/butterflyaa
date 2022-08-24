package command

import (
	// import local packages
	"github.com/impossible98/butterfly/internal/server"
)

func Serve() {
	server.InitServer()
}
