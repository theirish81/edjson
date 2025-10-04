package edjson

import (
	"github.com/jinzhu/copier"
)

// Converters is the package-level variable that holds all registered custom type converters.
// This variable is typically used by the Copy and MustCopy functions in this package.
var Converters TypeConverters

// TypeConverters is a custom type definition, essentially a slice of
// 'copier.TypeConverter' functions.
// Each TypeConverter function defines how to copy data between two fields
// of different types (e.g., converting a string field to a time.Time field).
type TypeConverters []copier.TypeConverter

// Add is a method on the TypeConverters type used to register one or more
// custom conversion functions to the collection.
//
// converters: A variadic argument of one or more 'copier.TypeConverter' functions.
func (t *TypeConverters) Add(converters ...copier.TypeConverter) {
	*t = append(*t, converters...)
}

// init Initializes the package-level Converters variable as an empty slice.
func init() {
	Converters = make(TypeConverters, 0)
}
