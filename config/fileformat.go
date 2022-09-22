package config

import (
	"encoding/xml"
	"log"
	"strings"
)

type Project struct {
	XMLName  xml.Name        `xml:"project"`
	Version  string          `xml:"version,attr"`
	Metadata ProjectMetadata `xml:"meta"`
	Build    ProjectBuild    `xml:"build"`
}

type ProjectBuild struct {
	XMLName      xml.Name            `xml:"build"`
	Type         string              `xml:"type,attr"`
	Dependencies ProjectDependencies `xml:"dependencies"`
	Steps        ProjectSteps        `xml:"steps"`
}

type ProjectDependencies struct {
	XMLName  xml.Name            `xml:"dependencies"`
	Contents []ProjectDependency `xml:"dependency"`
}

type ProjectDependency struct {
	XMLName xml.Name        `xml:"dependency"`
	Type    string          `xml:"type,attr"`
	Name    string          `xml:"name"`
	Version string          `xml:"version"`
	URL     string          `xml:"url"`
	Paths   []DepSearchPath `xml:"path"`
}

type DepSearchPath struct {
	XMLName xml.Name `xml:"path"`
	OS      string   `xml:"os,attr"`
	Path    string   `xml:",innerxml"`
}

type ProjectSteps struct {
	XMLName xml.Name `xml:"steps"`
	Step    []string `xml:"step"`
}

type ProjectMetadata struct {
	XMLName     xml.Name     `xml:"meta"`
	Name        string       `xml:"name"`
	Author      string       `xml:"author"`
	Description string       `xml:"description"`
	Kind        string       `xml:"kind"`
	Dates       ProjectDates `xml:"dates"`
}

type ProjectDates struct {
	XMLName xml.Name `xml:"dates"`
	Start   string   `xml:"start"`
	Update  string   `xml:"updated"`
}

type projKind string

// must be in lowercase
var ValidKinds = [...]projKind{
	"software",
}

func (k *projKind) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		s    string
		kind projKind
	)

	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	kind = projKind(strings.TrimSpace(strings.ToLower(s)))

	ispresent := false
	for _, b := range ValidKinds {
		if kind == b {
			ispresent = true
			break
		}
	}

	if !ispresent {
		log.Fatalf("invalid project kind %s; must be of %s", kind, ValidKinds)
	}

	*k = kind
	return nil
}
