package command

import (
	// import built-in packages
	"fmt"
	// import local packages
	"github.com/impossible98/butterfly/global"
)

// show the version of program
func Version() {
	fmt.Println(global.APP_NAME + " version: " + global.VERSION)
}
