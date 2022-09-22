package config

import (
	"fmt"
	"log"
	"strings"
)

func Confirm(prompt string) bool {
	var s string
	fmt.Printf("%s [y/n]? ", prompt)
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	s = strings.TrimSpace(strings.ToLower(s))

	return s[0] == 'y'
}
