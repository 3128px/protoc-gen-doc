package main_test

import (
	"bytes"
	"testing"

	. "github.com/3128px/protoc-gen-doc/v2/cmd/protoc-gen-doc"
	"github.com/stretchr/testify/require"
)

func TestHandleFlags(t *testing.T) {
	tests := []struct {
		args   []string
		result bool
	}{
		{[]string{"app", "-help"}, true},
		{[]string{"app", "-version"}, true},
		{[]string{"app", "-wjat"}, true},
	}

	for _, test := range tests {
		f := ParseFlags(new(bytes.Buffer), test.args)
		require.Equal(t, test.result, HandleFlags(f))
	}
}
