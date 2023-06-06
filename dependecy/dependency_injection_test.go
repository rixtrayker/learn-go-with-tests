package dependecy

import (
	"bytes"
	"testing"
)


func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Amr")

	got := buffer.String()
	want := "Hello, Amr"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}