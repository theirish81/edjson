package edjson

import "github.com/labstack/echo/v4"

// JSON is a helper function that performs a deep copy of a 'source' object
// into a new instance of type 'T', then sends the result as a JSON response
// using the Echo web framework.
// T: The desired type of the final JSON output struct.
// status: The HTTP status code to be returned (e.g., 200, 201).
// source: The initial data object to be copied from (e.g., a database model).
// err: An existing error that should immediately trigger a return if non-nil.
func JSON[T any](ctx echo.Context, status int, source any, err error) error {
	if err != nil {
		return err
	}
	out := new(T)

	// 3. Deep copy data from the 'source' object into the new 'out' object.
	// This step is  used to map a database model to a public Data Transfer Object (DTO).
	if err = Copy(out, source); err != nil {
		return err
	}
	return ctx.JSON(status, out)
}
