package rot13

import (
	"io"
)

type Rot13Reader struct {
	R io.Reader
}

const (
	a = 'a'
	z = 'z'
	A = 'A'
	Z = 'Z'
)

func Rot13(v byte) byte {
	// Returns true if min<=v<=max
	b := func(min, max, v byte) bool {
		return v >= min && v <= max
	}
	// Clamps a value between min and max
	c := func(min, max, v byte) byte {
		return min + (v-min)%(max-min+1)
	}

	if b(a, z, v) {
		return c(a, z, v+13)
	} else if b(A, Z, v) {
		return c(A, Z, v+13)
	}

	return v
}

func (t *Rot13Reader) Read(buf []byte) (int, error) {
	n, err := t.R.Read(buf)

	for i := range buf[:n] {
		buf[i] = Rot13(buf[i])
	}

	return n, err
}
