package commands

import (
	"encoding/xml"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"projectmgr/config"

	"github.com/fatih/color"
)

func Info() {
	var decodeddata config.Project

	title_ := color.New(color.FgBlue).Add(color.Underline)
	head2_ := color.New(color.FgGreen).Add(color.Underline)
	ok____ := color.New(color.FgGreen).Add(color.Bold)
	error_ := color.New(color.FgRed).Add(color.Bold)
	bold__ := color.New(color.Bold)
	norml_ := color.New(color.Reset)

	ProjectFileExists(func() {
		xmlfile, err := os.ReadFile(config.XML_FILENAME)
		if err != nil {
			log.Fatal(err)
		}
		err = xml.Unmarshal(xmlfile, &decodeddata)
		if err != nil {
			log.Fatal(err)
		}

		title_.Printf("project name\n")
		norml_.Printf("%s (%s)\n\n",
			decodeddata.Metadata.Name,
			decodeddata.Metadata.Kind,
		)

		title_.Printf("by\n")
		norml_.Printf("%s\n\n",
			decodeddata.Metadata.Author,
		)

		title_.Printf("description\n")
		norml_.Printf("%s\n\n",
			config.Strip(decodeddata.Metadata.Description),
		)

		title_.Printf("project dates\n")
		bold__.Printf("started on:   ")
		norml_.Printf("%s\n",
			decodeddata.Metadata.Dates.Start,
		)
		bold__.Printf("last updated: ")
		norml_.Printf("%s\n\n",
			decodeddata.Metadata.Dates.Update,
		)

		title_.Printf("build dependencies\n")
		for _, builddep := range decodeddata.Build.Dependencies.Contents {
			deptype := builddep.Type

			bold__.Printf("%s", builddep.Name)
			if strings.TrimSpace(builddep.Version) != "" {
				bold__.Printf(" %s", builddep.Version)
			}
			norml_.Printf(" (%s)", deptype)
			if strings.TrimSpace(builddep.URL) != "" {
				bold__.Printf(" <%s>", builddep.URL)
			}

			norml_.Printf("\n")

			searchPathFound := func(entry string, deptype string) bool {
				if strings.TrimSpace(deptype) == "software" {
					_, err := exec.LookPath(entry)
					return err == nil
				} else {
					_, err := os.Stat(entry)
					return err == nil
				}
			}

			isSupportedPath := func(searchpath config.DepSearchPath) bool {
				return strings.TrimSpace(searchpath.OS) == "" || strings.TrimSpace(searchpath.OS) == runtime.GOOS
			}

			filter := func(vs []config.DepSearchPath, f func(config.DepSearchPath) bool) []config.DepSearchPath {
				filtered := make([]config.DepSearchPath, 0)
				for _, v := range vs {
					if f(v) {
						filtered = append(filtered, v)
					}
				}
				return filtered
			}

			// get search paths that are common OR belong to the current OS
			filteredpaths := filter(builddep.Paths, func(entry config.DepSearchPath) bool {
				return isSupportedPath(entry)
			})

			if len(filteredpaths) > 0 {
				norml_.Printf("  ")
				head2_.Printf("search paths\n")
				for _, searchpath := range builddep.Paths {
					if isSupportedPath(searchpath) {
						norml_.Printf("    - ")
						if searchPathFound(searchpath.Path, deptype) {
							ok____.Printf("[FOUND]")
						} else {
							error_.Printf("[MISS!]")
						}
						norml_.Printf(" %s\n", searchpath.Path)
					}
				}
				norml_.Printf("\n")
			}
		}

		title_.Printf("build steps\n")
		for _, buildstep := range decodeddata.Build.Steps.Step {
			norml_.Printf("  - %s\n", config.Strip(buildstep))
		}
		// TODO
	})
}
