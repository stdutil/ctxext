# ctxext

`ctxext` is a utility package for safely setting and retrieving values in `context.Context` within a Go application.

> [!WARNING]
> This package must be referenced directly in your application and should not be included in a publicly published package.

## Author

Elizalde Baguinon
Created: May 25, 2025

## Features

- **GetValue**: Retrieve a value of the specified type from the context.
- **SetValue**: Store a key-value pair in a new context.
- **HasValue**: Check whether a key exists in the context.

## Installation

To use ctxext, import it into your Go project:

```go
import "github.com/stdutil/ctxext"
```

## Usage

### Defining Context Keys

```go
package ctxext

type ContextKey string

const (
    UserName  ContextKey = "user_name"
    SessionID ContextKey = "session_id"
)
```

### Setting and Retrieving Context Values

```go
package main

import (
    "context"
    "fmt"
    "github.com/stdutil/ctxext"
)

func main() {
    ctx := context.Background()

    // Store a value in the context
    ctx = ctxext.SetValue(ctx, ctxext.UserName, "Alice")

    // Retrieve the value
    username := ctxext.GetValue[string](ctx, ctxext.UserName)
    fmt.Println("Username:", username) // Output: Username: Alice

    // Check if a key exists
    exists := ctxext.HasValue(ctx, ctxext.UserName)
    fmt.Println("Key exists:", exists) // Output: Key exists: true
}
```

## Why Use ctxext?

- Avoid type collisions with built-in `context.Context` storage
- Type-safe retrieval using generics
- Simple and efficient utility functions

## License

This project is licensed under the MIT License.

---

Feel free to modify this README to match your project details, such as updating the repository link!
