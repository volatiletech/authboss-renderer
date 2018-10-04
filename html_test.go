package abrenderer

import (
	"context"
	"strings"
	"testing"
)

func TestRenderer(t *testing.T) {
	r := NewHTML("/auth", ".")

	err := r.Load("login")
	if err != nil {
		t.Error(err)
	}

	o, content, err := r.Render(context.Background(), "login", nil)
	if err != nil {
		t.Fatal(err)
	}

	if content != "text/html" {
		t.Error("context type not set properly")
	}

	if len(o) == 0 {
		t.Error("it should have rendered a template")
	}

	if !strings.Contains(string(o), "/auth/login") {
		t.Error("expected the url to be rendered out for the form post location")
	}
}
