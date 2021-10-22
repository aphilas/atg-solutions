package rot13_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/nevilleomangi/atg-solutions/rot13"
)

func TestRot13Reader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13.Rot13Reader{s}

	buf := &bytes.Buffer{}
	io.Copy(buf, &r)
	got := buf.String()

	want := "You cracked the code!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
