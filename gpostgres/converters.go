package gpostgres

import (
	"github.com/jinzhu/copier"
	// Imports the GORM datatypes package, which includes the datatypes.JSONType.
	"gorm.io/datatypes"
)

// GenericJsonTypeConverter creates converters for copying data **between** a
// gorm.io/datatypes.JSONType[T] wrapper and its underlying Go type T.
// This is used when the source and destination types are the same structure,
// only one is wrapped in the JSONType.
//
// T: The underlying Go struct type held within the JSONType wrapper.
// Returns: A slice of two `copier.TypeConverter` structs.
func GenericJsonTypeConverter[T any]() []copier.TypeConverter {
	return []copier.TypeConverter{
		// Converter 1: JSONType[T] -> T (Unwrap)
		{
			SrcType: *new(datatypes.JSONType[T]),
			DstType: *new(T),
			Fn: func(src interface{}) (dst interface{}, err error) {
				t := src.(datatypes.JSONType[T]).Data()
				return t, err
			},
		},
		// Converter 2: T -> JSONType[T] (Wrap)
		{
			DstType: *new(datatypes.JSONType[T]),
			Fn: func(src interface{}) (dst interface{}, err error) {
				t := datatypes.NewJSONType(src.(T))
				return t, err
			},
		},
	}
}

// JsonTypeConverter creates converters for copying data **between** a
// gorm.io/datatypes.JSONType[S] wrapper and a *different* Go type T.
// This requires an inner `copier.Copy` call because the source and destination
// structs (S and T) may have different fields.
//
// S: The underlying Go struct type held within the JSONType wrapper.
// T: The plain Go struct type being copied to/from.
// Returns: A slice of two `copier.TypeConverter` structs.
func JsonTypeConverter[S any, T any]() []copier.TypeConverter {
	return []copier.TypeConverter{
		// Converter 1: JSONType[S] -> T (Unwrap and Copy)
		{
			SrcType: *new(datatypes.JSONType[S]),
			DstType: *new(T),
			Fn: func(src interface{}) (dst interface{}, err error) {
				t := new(T)
				err = copier.Copy(t, src.(datatypes.JSONType[S]).Data())
				return *t, err
			},
		},
		// Converter 2: T -> JSONType[S] (Copy and Wrap)
		{
			SrcType: *new(T),
			DstType: *new(datatypes.JSONType[S]),
			Fn: func(src interface{}) (dst interface{}, err error) {
				s := new(S)
				err = copier.Copy(s, src.(T))
				return datatypes.NewJSONType(*s), err
			},
		},
	}
}
