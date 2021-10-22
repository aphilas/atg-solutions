package ip_test

import (
	"fmt"
	"testing"

	"github.com/nevilleomangi/atg-solutions/ip"
)

func TestIPStringer(t *testing.T) {
	cases := []struct {
		host string
		ip   ip.IPAddr
		want string
	}{
		{"loopback", ip.IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"googleDNS", ip.IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}

	for _, c := range cases {
		got := fmt.Sprint(c.ip)
		if got != c.want {
			t.Fatalf("%q: got %q, want %q", c.host, got, c.want)
		}
	}
}
