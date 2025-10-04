# EDJSON: Echo DTO JSON ✨

EDJSON is a small, utility-focused Go library that simplifies **data mapping** and **JSON response generation** within the **Echo** web framework. It builds upon robust **deep-copy** functionality to ensure clean, reliable data transfer from backend models to public-facing **Data Transfer Objects (DTOs)**.

Its primary goal is to provide a single, clean function for mapping complex data structures into a defined response format before delivery.

-----

## 🚀 Main Feature: The `JSON` Helper

The core of EDJSON is the generic `JSON` function. It is designed to eliminate boilerplate code in your HTTP handlers by combining error checking, deep-copying, DTO instantiation, and JSON response delivery into one clean call.

### How It Works

The `JSON` function handles three essential steps for you:

1.  **Error Check:** Immediately returns if a previous error (e.g., from a data service call) is passed to it.
2.  **DTO Mapping:** **Deep-copies** the data from your `source` object into a **newly created** instance of the specified DTO type `T`.
3.  **JSON Response:** Sends the resulting DTO as a JSON response using the Echo context with the provided HTTP `status`.

### Usage Example

```go
// T is the DTO type (e.g., UserResponseDTO)
// source is the internal data structure (e.g., *User)
func GetUser(c echo.Context) error {
    user, dataServiceErr := service.Users.Get(c.Param("id"))
    
    // Maps 'user' to UserResponseDTO, handles dataServiceErr, and sends the JSON response.
    return edjson.JSON[UserResponseDTO](c, http.StatusOK, user, dataServiceErr)
}
```

-----

## 🛠️ Data Copying & Custom Conversion

EDJSON ensures all data copying is performed as a **deep copy**, meaning nested structs, maps, and slices are copied by value, preventing unexpected side effects or mutation of your source data.

### Custom Converters

EDJSON manages a global list of **Type Converters** which allow you to define custom logic for copying between fields of different types. This is essential for:

* **Custom Types:** Handling conversions between library-specific types and standard Go structs.
* **Encoding/Decoding:** Automatic encoding/decoding during the copy process (e.g., converting a `[]byte` field to a Base64-encoded `string`).

You can easily register your custom converters during application startup:

```go
func init() {
    // Add built-in converters for []byte -> Base64 string
    edjson.Converters.Add(edjson.BytesToBase64Converter()...)
}
```

-----