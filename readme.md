# ctxext

ctxext is a lightweight utility package for working with `context.Context` in Go. It provides convenient functions for safely setting, retrieving, and checking values within a context.

> [!WARNING]
> This package must be referenced directly in your application. Do not put it in a package that you intend to publish.


## Features

- **GetValue**: Retrieve a value of the specified type from the context.
- **SetValue**: Store a key-value pair in a new context.
- **HasValue**: Check whether a key exists in the context.

## Installation

To use ctxext, simply import it into your Go project:

```go
import "github.com/stdutil/ctxext"
```

## Usage

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
    ctx = ctxext.SetValue(ctx, "username", "Alice")

    // Retrieve the value
    username := ctxext.GetValue[string](ctx, "username")
    fmt.Println("Username:", username) // Output: Username: Alice

    // Check if a key exists
    exists := ctxext.HasValue(ctx, "username")
    fmt.Println("Key exists:", exists) // Output: Key exists: true
}
```

## Why Use ctxext?

- Avoid type collisions with `context.WithValue`
- Type-safe retrieval using generics
- Simple and efficient utility functions

## License

This project is licensed under the MIT License.

---

Feel free to modify this README to match your project details, such as updating the repository link!
