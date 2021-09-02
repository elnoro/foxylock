package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	h, err := newHostsFileFromContents("127.0.0.1    test.example # added by foxylock")
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1    test.example # added by foxylock", h.toString())
}

// found by fuzz
func TestToStringInvalid(t *testing.T) {
	t.Skip("not everything is fixed")
	cases := []struct {
		expected string
	}{
		{"# added by foxylock"},
		{"1 e # added by foxylock"},
	}

	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			h, err := newHostsFileFromContents(c.expected)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, h.toString())
		})
	}
}
