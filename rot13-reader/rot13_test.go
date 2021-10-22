package rot13reader_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	reader "github.com/nevilleomangi/atg-solutions/rot13-reader"
)

func TestRot13Reader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := reader.Rot13Reader{s}

	buf := &bytes.Buffer{}
	io.Copy(buf, &r)
	got := buf.String()

	want := "You cracked the code!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
