package whois

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name              string
		expectedAvailable bool
	}{
		{"140bedbf9c3f6d56a9846d.br", true},
		{"140bedbf9c3f6d56a9846d2ba7088798683f4da0c248231336e6a05679e4fdfe.com", true},
		{"example.co.uk", false},
		{"example.com", false},
		{"google.io", false},
	}

	for _, test := range tests {
		filename := fmt.Sprintf("testdata/whois_%s.txt", test.name)
		content, err := ioutil.ReadFile(filename)
		assert.NoError(t, err, "cannot read test file %s", filename)
		if err != nil {
			return
		}

		attr := Parse(test.name, string(content))
		assert.Equalf(t, test.expectedAvailable, attr.Available, "Wrong available value for %s", test.name)
	}
}
