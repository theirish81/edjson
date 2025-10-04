package edjson

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCopyBytes struct {
	Data []byte
}
type TestCopyString struct {
	Data string
}

func TestCopy(t *testing.T) {
	Converters.Add(BytesToBase64Converter()...)
	out := MustCopy[TestCopyString](TestCopyBytes{
		Data: []byte("test"),
	})
	testOut, _ := base64.StdEncoding.DecodeString(out.Data)
	assert.Equal(t, testOut, []byte("test"))
}
