package input

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Stdin() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if (info.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		return "", fmt.Errorf("Not a valid input %d", 1)
	} else if info.Size() > 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		return string(data), err
	}

	return "", nil
}
