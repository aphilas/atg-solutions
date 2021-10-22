package main

import (
	"fmt"
	"testing"
)

func TestIPStringer(t *testing.T) {
	cases := []struct {
		host string
		ip   IPAddr
		want string
	}{
		{"loopback", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"googleDNS", IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}

	for _, c := range cases {
		got := fmt.Sprint(c.ip)
		if got != c.want {
			t.Fatalf("%q: got %q, want %q", c.host, got, c.want)
		}
	}
}
