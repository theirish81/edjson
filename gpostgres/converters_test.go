package gpostgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theirish81/edjson"
	"gorm.io/datatypes"
)

type s1 struct {
	Data string
}
type s2 struct {
	Data string
}

func TestGenericJsonTypeConverter(t *testing.T) {
	edjson.Converters.Add(GenericJsonTypeConverter[s1]()...)
	sx := edjson.MustCopy[s1](datatypes.NewJSONType(s1{
		Data: "test",
	}))
	assert.Equal(t, sx.Data, "test")
}

func TestJSONTypeConverter(t *testing.T) {
	edjson.Converters.Add(JsonTypeConverter[s1, s2]()...)
	sx := edjson.MustCopy[s2](datatypes.NewJSONType(s1{
		Data: "test",
	}))
	assert.Equal(t, sx.Data, "test")
}
