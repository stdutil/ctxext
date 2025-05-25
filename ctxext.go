package ctxext

import "context"

type contextKey string

// GetValue retrieves a context value of the specified type.
//
// The key is converted to an internal contextKey, and the associated value is returned.
// If the key is not found or the value is of the wrong type, the zero value is returned.
func GetValue[T any](ctx context.Context, key string) T {
	var zero T
	v := ctx.Value(contextKey(key))
	if v == nil {
		return zero
	}
	rv, ok := v.(T)
	if !ok {
		return zero
	}
	return rv
}

// SetValue returns a new context with the given key and value set.
//
// The key is converted to an internal contextKey to avoid collisions.
func SetValue[T any](ctx context.Context, key string, value T) context.Context {
	return context.WithValue(ctx, contextKey(key), value)
}

// HasValue checks whether the context contains a non-nil value for the given key.
func HasValue(ctx context.Context, key string) bool {
	return ctx.Value(contextKey(key)) != nil
}
