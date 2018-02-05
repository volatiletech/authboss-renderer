// Package abrenderer implements a basic
// html/template renderer for an app.
package abrenderer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/volatiletech/authboss"
)

//go:generate go-bindata -prefix templates -pkg abrenderer ./templates

// Renderer for authboss
type Renderer struct {
	mountpath string
	layout    *template.Template
	templates map[string]*template.Template
}

// New renderer
func New(mountpath string) *Renderer {
	r := &Renderer{
		mountpath: mountpath,
		templates: make(map[string]*template.Template),
	}

	return r
}

// Load a template
func (r *Renderer) Load(names ...string) error {
	if r.layout == nil {
		b, err := Asset("layout.html.tpl")
		if err != nil {
			return err
		}

		r.layout, err = template.New("").Parse(string(b))
		if err != nil {
			return errors.Wrap(err, "failed to load layout template")
		}
	}

	for _, n := range names {
		filename := fmt.Sprintf("%s.html.tpl", n)
		b, err := Asset(filename)
		if err != nil {
			return err
		}

		clone, err := r.layout.Clone()
		if err != nil {
			return err
		}

		funcmap := template.FuncMap{
			"title": strings.Title,
			"mountpathed": func(location string) string {
				if r.mountpath == "/" {
					return location
				}
				return path.Join(r.mountpath, location)
			},
		}

		_, err = clone.New("authboss").Funcs(funcmap).Parse(string(b))
		if err != nil {
			return errors.Wrapf(err, "failed to load template for page %s", n)
		}

		r.templates[n] = clone
	}

	return nil
}

// Render a view
func (r *Renderer) Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	buf := &bytes.Buffer{}

	tpl, ok := r.templates[page]
	if !ok {
		return nil, "", errors.Errorf("template for page %s not found", page)
	}

	err = tpl.Execute(buf, data)
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed to render template for page %s", page)
	}

	return buf.Bytes(), "text/html", nil
}
