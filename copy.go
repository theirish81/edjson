package edjson

import "github.com/jinzhu/copier"

// Copy performs a **deep copy** of fields from the 'source' object to the 'target' object.
// It uses the 'copier' library with specific options for deep copying and custom conversions.
//
// target: A pointer to the destination (struct, map, etc.) where data will be copied into.
// source: The source object (struct, map, slice, etc.) to copy data from.
// Returns: An error if the copy operation fails.
func Copy(target any, source any) error {
	return copier.CopyWithOption(target, source, copier.Option{
		DeepCopy:   true,
		Converters: Converters},
	)
}

// MustCopy is a generic helper that performs a **deep copy** into a **newly created**
// instance of the specified type 'T', ignoring any potential errors.
//
// T: The value type of the destination object to be created and returned.
// source: The object to copy data from.
// Returns: A new instance of type 'T' containing the copied data.
// NOTE: Use with caution, as ignoring the error means a failed copy will result in
// an uninitialized or partially-initialized object.
func MustCopy[T any](source any) T {
	target := new(T)
	_ = Copy(target, source)
	return *target
}
