package edjson

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockResponseWriter struct {
	httptest.ResponseRecorder
	BodyContent []byte
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	m.BodyContent = append(m.BodyContent, b...)
	return m.ResponseRecorder.Write(b)
}

type s1 struct {
	Data string
}
type s2 struct {
	Data string
}
type a struct {
	TheString string
	TheStruct s1
	Binary    []byte
}

type b struct {
	TheString string
	TheStruct s2
	Binary    string
}

func TestJSON(t *testing.T) {
	Converters.Add(BytesToBase64Converter()...)

	req := httptest.NewRequest(http.MethodGet, "/foo", nil)
	rw := &MockResponseWriter{
		ResponseRecorder: *httptest.NewRecorder(),
	}
	e := echo.New()
	c := e.NewContext(req, rw)

	err := JSON[b](c, 200, a{
		TheString: "test",
		Binary:    []byte("test"),
		TheStruct: s1{
			Data: "test",
		},
	}, nil)
	assert.Nil(t, err)
	var out b
	assert.Nil(t, json.Unmarshal(rw.BodyContent, &out))
	assert.Equal(t, out.TheString, "test")
	assert.Equal(t, out.Binary, base64.StdEncoding.EncodeToString([]byte("test")))
}
