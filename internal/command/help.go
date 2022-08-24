package command

import (
	// import built-in packages
	"fmt"
	"strings"
	// import local packages
	"github.com/impossible98/butterfly/global"
)

var serverHelp = "\tserve\t\tStart the server\n"
var versionHelp = "\t--version\tShow version\n"

func Help() {
	fmt.Println("\n" + strings.ToLower(global.APP_NAME) + " help:\n\n" + serverHelp + versionHelp)
}
