package commands

import (
	"fmt"
	"os"
	"projectmgr/config"
)

func ProjectFileExists(exec func()) {
	if _, err := os.Stat(config.XML_FILENAME); err == nil { // exists
		exec()
	} else {
		notexistprompt := config.Confirm("no project file exists, create one")
		if notexistprompt {
			Init()
		} else {
			fmt.Println("not creating...")
		}

	}
}
