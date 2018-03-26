// Package abrenderer implements a basic template renderer for a
// web app for both html responses and both or either html/txt
// e-mail rendering.
package abrenderer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:generate go-bindata -pkg abrenderer ./html-templates ./email-templates

func loadWithOverride(override string, name string) ([]byte, error) {
	if len(override) != 0 {
		file := filepath.Join(override, name)

		b, err := ioutil.ReadFile(file)
		if err == nil {
			return b, err
		} else if os.IsNotExist(err) {
			// Fall through
		} else {
			return nil, err
		}
	}

	return Asset(name)
}
