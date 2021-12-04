package reader

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) (string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contents)), nil
}
