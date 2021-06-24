package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassFail(t *testing.T) {
	tests := []struct {
		name string
		pass bool
	}{
		{
			name: "given pass=true, " +
				"when assert.True is invoked, " +
				"then it always passes",
			pass: true,
		},
		{
			name: "given pass=false, " +
				"when assert.True is invoked, " +
				"then it always fails",
			pass: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, tt.pass, "test doesn't pass")
		})
	}
}
