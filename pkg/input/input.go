package input

import (
	"strings"
)

//GrabDomainNames parses domain names from command line input string, It
//converts names to lower case in the process.
func GrabDomainNames(s string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(strings.Replace(s, ",", " ", -1))))
}

//ValidateDomainNames does nothing for now
func ValidateDomainNames(names ...string) error {
	return nil
}
