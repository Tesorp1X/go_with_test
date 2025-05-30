package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Liz")

	got := buffer.String()
	want := "Hello, Liz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
