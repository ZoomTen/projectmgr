package commands

import (
	"log"
	"path/filepath"
	"projectmgr/config"

	"github.com/pkg/browser"
)

func Web() {
	ProjectFileExists(func() {
		afp, err := filepath.Abs(config.XML_FILENAME)
		if err != nil {
			log.Fatal(err)
		}
		err = browser.OpenURL("file://" + afp)
		if err != nil {
			log.Fatal(err)
		}
	})
}
