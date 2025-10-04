# EDJSON GPOSTGRES: GORM Datatype Converters for EDJSON

This package provides essential **custom type converters** specifically designed for use with the **EDJSON** core library and the GORM database framework, particularly when dealing with **JSON/JSONB** column types.

It enables EDJSON's deep-copy functionality (`edjson.Copy` and `edjson.JSON`) to seamlessly map data between the GORM wrapper type (`datatypes.JSONType[T]`) and your standard Go structs.

## 📦 Usage

To enable these conversions, you must register them with the core EDJSON package's global converter list, typically in an `init()` function or application setup:

```go
import (
    "github.com/theirish81/edjson"
    "github.com/theirish81/edjson/gpostgres" // Assuming this is your package path
    "gorm.io/datatypes"
)

// Define your structs
type Settings struct { /* ... */ }
type DTO struct { /* ... */ }

func init() {
    // Register the converters globally in EDJSON
    edjson.Converters.Add(
        // For mapping between datatypes.JSONType[Settings] <-> Settings
        edjson_gorm.GenericJsonTypeConverter[Settings]()...,
        
        // For mapping between datatypes.JSONType[Settings] <-> DTO (different structs)
        edjson_gorm.JsonTypeConverter[Settings, DTO]()...,
    )
}
```

## ✨ Converter Functions

This package provides two generic functions to generate the required bidirectional converters:

### 1\. `GenericJsonTypeConverter[T]`

This function creates converters for mapping **between a GORM JSON type and its identical underlying Go struct**.

| Direction | Purpose |
| :--- | :--- |
| `datatypes.JSONType[T]` → `T` | **Unwraps** the data from the GORM JSON container to the plain Go struct. |
| `T` → `datatypes.JSONType[T]` | **Wraps** the plain Go struct into the GORM JSON container. |

### 2\. `JsonTypeConverter[S, T]`

This function creates converters for complex mapping **between a GORM JSON type (based on struct S) and a different destination struct (T)**.

This is critical when you need to map a database JSON field (which holds struct `S`) to a public DTO (`T`). This involves an **inner copy operation** during the conversion process to handle field differences.

| Direction | Process                                                                                                   |
| :--- |:----------------------------------------------------------------------------------------------------------|
| `datatypes.JSONType[S]` → `T` | Unwraps struct `S`, then uses `edjson.Copy` to map fields from `S` to `T`.                                |
| `T` → `datatypes.JSONType[S]` | Uses `edjson.Copy` to map fields from `T` to a new struct `S`, then wraps `S` in the GORM JSON container. |