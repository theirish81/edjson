package edjson

import (
	"encoding/base64"

	"github.com/jinzhu/copier"
)

// BytesToBase64Converter creates and returns a slice of custom type converters
// for the `copier` library.
// These converters handle the automatic conversion of byte slices ([]byte or []uint8)
// to their Base64-encoded string representation during a copy operation.
//
// Returns: A slice of `copier.TypeConverter` structs.
func BytesToBase64Converter() []copier.TypeConverter {
	return []copier.TypeConverter{
		// Converter 1: Converts a standard Go byte slice ([]byte) to a Base64 string.
		{
			SrcType: []byte{}, // Source type is a byte slice.
			DstType: "",       // Destination type is a string.
			Fn: func(src interface{}) (dst interface{}, err error) {
				// Encodes the source []byte slice using standard Base64 encoding.
				return base64.StdEncoding.EncodeToString(src.([]byte)), nil
			},
		},
		// Converter 2: Converts an alias for byte slice ([]uint8) to a Base64 string.
		{
			SrcType: []uint8{}, // Source type is a slice of uint8 (identical to []byte).
			DstType: "",        // Destination type is a string.
			Fn: func(src interface{}) (dst interface{}, err error) {
				// Encodes the source []uint8 slice using standard Base64 encoding.
				return base64.StdEncoding.EncodeToString(src.([]uint8)), nil
			},
		},
	}
}
