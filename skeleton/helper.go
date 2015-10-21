package skeleton

import (
	"bytes"
	"os"
	"regexp"
)

var reg = regexp.MustCompile("[A-z0-9]*")

// camelCase transform string to CamelCase
func camelCase(s string) string {
	b := []byte(s)
	matched := reg.FindAll(b, -1)
	for i, m := range matched {
		if i == 0 {
			continue
		}
		matched[i] = bytes.Title(m)
	}

	return string(bytes.Join(matched, nil))
}

// mkdir makes the named directory.
func mkdir(dir string) error {
	if _, err := os.Stat(dir); err == nil {
		return nil
	}

	return os.MkdirAll(dir, 0777)
}
