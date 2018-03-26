// Package abrenderer implements a basic
// html/template renderer for an app.
package abrenderer

import (
	"bytes"
	"context"
	"fmt"
	htmltemplate "html/template"
	txttemplate "html/template"
	"io"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/volatiletech/authboss"
)

// Email renderer for authboss, renders using html/template
// Allows overrides of the same template names in the same prefixes.
// For example:
// If overridePath is /home/authboss/views
// You could override the confirm_html.tpl by creating a file at
// /home/authboss/views/email-templates/confirm.html.tpl
//
// This renderer differentiates html and text templates
// by checking for a suffix _html or _txt in the filename and
// uses the appropriate template library for the file type.
type Email struct {
	mountPath    string
	overridePath string

	htmlTemplates map[string]*htmltemplate.Template
	txtTemplates  map[string]*txttemplate.Template

	funcMap map[string]interface{}
}

// NewEmail renderer
func NewEmail(mountPath, overridePath string) *Email {
	e := &Email{
		mountPath:     mountPath,
		overridePath:  overridePath,
		htmlTemplates: make(map[string]*htmltemplate.Template),
		txtTemplates:  make(map[string]*txttemplate.Template),
		funcMap: txttemplate.FuncMap{
			"title": strings.Title,
			"mountpathed": func(location string) string {
				if mountPath == "/" {
					return location
				}
				return path.Join(mountPath, location)
			},
		},
	}

	return e
}

// Load a template
func (e *Email) Load(names ...string) error {
	for _, n := range names {
		filename := fmt.Sprintf("email-templates/%s.tpl", n)
		b, err := loadWithOverride(e.overridePath, filename)
		if err != nil {
			return errors.Wrapf(err, "failed to load template %s", filename)
		}

		if strings.HasSuffix(n, "_txt") {
			txt, err := txttemplate.New("authboss").Funcs(e.funcMap).Parse(string(b))
			if err != nil {
				return errors.Wrapf(err, "failed to load txt template for page %s", n)
			}
			e.txtTemplates[n] = txt
		} else {
			html, err := htmltemplate.New("authboss").Funcs(e.funcMap).Parse(string(b))
			if err != nil {
				return errors.Wrapf(err, "failed to load html template for page %s", n)
			}
			e.htmlTemplates[n] = html
		}

	}

	return nil
}

// Render a view
func (e *Email) Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	buf := &bytes.Buffer{}

	var exe executor
	var ok bool
	if strings.HasSuffix(page, "_txt") {
		exe, ok = e.txtTemplates[page]
		contentType = "text/plain"
	} else {
		exe, ok = e.htmlTemplates[page]
		contentType = "text/html"
	}

	if !ok {
		return nil, "", errors.Errorf("template for page %s not found", page)
	}

	err = exe.Execute(buf, data)
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed to render template for page %s", page)
	}

	return buf.Bytes(), contentType, nil
}

type executor interface {
	Execute(io.Writer, interface{}) error
}
