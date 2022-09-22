package commands

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"projectmgr/config"
)

//go:embed "files/project.zumi.xml"
var initxml string

func Init() {
	var (
		err              error
		_                int
		f                *os.File
		should_overwrite bool
	)
	if _, err = os.Stat(config.XML_FILENAME); err == nil {
		should_overwrite = config.Confirm("project file exists, overwrite")
		if !should_overwrite {
			fmt.Println("not overwriting.")
			return
		} else {
			fmt.Println("project file overwritten")
		}
	} else {
		fmt.Println("creating project file")
	}
	f, err = os.Create(config.XML_FILENAME) // hard coded
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(initxml)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("open the project file (%s) for further details...", config.XML_FILENAME)
}
