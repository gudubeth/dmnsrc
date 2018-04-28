package checker

import (
	"github.com/domainr/whois"
)

// Whois gets whois information of name using whoisServer.
//
// It returns whois data and any error encountered.
func Whois(name string) (string, error) {
	request, err := whois.NewRequest(name)
	if err != nil {
		return "", err
	}
	var client = whois.NewClient(whois.DefaultTimeout)
	response, err := client.Fetch(request)
	return response.String(), err
}
