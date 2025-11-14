package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"ADD 1 2", Add{Val{1}, Val{2}}},
		{"SUB 1 2", Sub{Val{1}, Val{2}}},
		{"MUL 3 4", Mul{Val{3}, Val{4}}},
		{"DIV 10 2", Div{Val{10}, Val{2}}},
		{"MOD 10 3", Mod{Val{10}, Val{3}}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			res := Parse(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
