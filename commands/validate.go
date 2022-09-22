package commands

import (
	_ "embed"
	"fmt"
	//"github.com/lestrrat-go/libxml2"
	//"github.com/lestrrat-go/libxml2/xsd"
)

//go:embed "files/project.zumi.xsd"
var xsdsrc []byte

func Validate() {
	ProjectFileExists(func() {
		fmt.Println("TODO")
		//schema, err := xsd.Parse(xsdsrc)
		//if err != nil {
		//	panic(err)
		//}
	})
}
