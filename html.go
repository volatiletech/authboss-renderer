package abrenderer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"path"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/authboss/v3"
)

// HTML renderer for authboss, renders using html/template
// Allows overrides of the same template names in the same prefixes.
// For example:
// If overridePath is /home/authboss/views
// You could override the login.tpl by creating a file at
// /home/authboss/views/html-templates/login.tpl
type HTML struct {
	mountPath    string
	overridePath string

	layout    *template.Template
	templates map[string]*template.Template

	funcMap map[string]interface{}
}

// NewHTML renderer
func NewHTML(mountPath string, overridePath string) *HTML {
	h := &HTML{
		mountPath:    mountPath,
		overridePath: overridePath,
		templates:    make(map[string]*template.Template),

		funcMap: template.FuncMap{
			"title": strings.Title,
			"mountpathed": func(location string) string {
				if mountPath == "/" {
					return location
				}
				return path.Join(mountPath, location)
			},
		},
	}

	return h
}

// Load a template
func (h *HTML) Load(names ...string) error {
	if h.layout == nil {
		b, err := loadWithOverride(h.overridePath, "html-templates/layout.tpl")
		if err != nil {
			return err
		}

		h.layout, err = template.New("").Funcs(h.funcMap).Parse(string(b))
		if err != nil {
			return errors.Wrap(err, "failed to load layout template")
		}
	}

	for _, n := range names {
		filename := fmt.Sprintf("html-templates/%s.tpl", n)
		b, err := loadWithOverride(h.overridePath, filename)
		if err != nil {
			return err
		}

		clone, err := h.layout.Clone()
		if err != nil {
			return err
		}

		_, err = clone.New("authboss").Funcs(h.funcMap).Parse(string(b))
		if err != nil {
			return errors.Wrapf(err, "failed to load template for page %s", n)
		}

		h.templates[n] = clone
	}

	return nil
}

// Render a view
func (h *HTML) Render(_ context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	buf := &bytes.Buffer{}

	tpl, ok := h.templates[page]
	if !ok {
		return nil, "", errors.Errorf("template for page %s not found", page)
	}

	err = tpl.Execute(buf, data)
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed to render template for page %s", page)
	}

	return buf.Bytes(), "text/html", nil
}
