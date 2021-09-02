//go:build gofuzzbeta
// +build gofuzzbeta

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzHostsfileToString(f *testing.F) {
	f.Add("127.0.0.1	test.test")
	f.Add("1.1.1.1    test.example # foxylock")
	// f.Add("127.0.0.1    test.example # added by foxylock")
	f.Fuzz(func(t *testing.T, hosts string) {
		h, err := newHostsFileFromContents(hosts)
		if err != nil {
			t.Fatalf("Failed to parse")
		}

		if h.toString() != hosts {
			assert.Equal(t, hosts, h.toString())
		}
	})
}
