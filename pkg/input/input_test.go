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
	}

	for _, test := range tests {
		res := GrabDomainNames(test.in)
		assert.Equalf(t, test.expected, res, "failed %s", test.desc)
	}
}
