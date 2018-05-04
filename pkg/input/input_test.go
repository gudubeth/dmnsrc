package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrabDomainsNames(t *testing.T) {
	tests := []struct {
		in       string
		expected []string
		desc     string
	}{
		{"", []string{}, "parsing empty string"},
		{"    ", []string{}, "parsing string consist of spaces"},
		{" ,,,   ", []string{}, "parsing string consist of spaces and commas"},
		{"example.com", []string{"example.com"}, "parsing single domain"},
		{"example.com example.org", []string{"example.com", "example.org"}, "parsing domains seperated with space"},
		{"example.com,example.org", []string{"example.com", "example.org"}, "parsing domains seperated with comma"},
		{",example.com,example.org", []string{"example.com", "example.org"}, "parsing domains prefixed with comma"},
		{"example.com,example.org,", []string{"example.com", "example.org"}, "parsing domains suffixed with comma"},
		{" example.com   example.org     ", []string{"example.com", "example.org"}, "parsing domains with spaces everywhere"},
		{" , , example.com ,  ,,,,example.org  ,,,,  ,, ", []string{"example.com", "example.org"}, "parsing domains with spaces and commas everywhere"},
		{"\t\texample.com   \n\texample.org     ", []string{"example.com", "example.org"}, "parsing domains with tabs and new lines"},
		{"\t\texample.com   \n\texample.org     \r\nexample.net\r\n", []string{"example.com", "example.org", "example.net"}, "parsing domains with tabs and new lines"},
	}

	for _, test := range tests {
		res := GrabDomainNames(test.in)
		assert.Equalf(t, test.expected, res, "failed %s", test.desc)
	}
}
