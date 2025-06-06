// ctxext is a package to set and get values from a context
//
// This package must be referenced directly in your application.
// Do not put it in a package that you intend to publish.
//
// Author: Elizalde Baguinon
// Created: May 25, 2025
package ctxext

import "context"

// ContextKey is a type to make a key to avoid collisions
type ContextKey string

const (
	UserName      ContextKey = "user_name"
	SessionID     ContextKey = "session_id"
	CacheDuration ContextKey = "cache_duration"
	SecretKeys    ContextKey = "secret_keys"
	MaxRecords    ContextKey = "max_records"
	PageSize      ContextKey = "page_size"
)

// GetValue retrieves a context value of the specified type.
//
// If the key is not found or the value is of the wrong type, the zero value is returned.
func GetValue[T any](ctx context.Context, key ContextKey) T {
	var zero T
	v := ctx.Value(key)
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
// The key is a built-in type to avoid collisions.
func SetValue[T any](ctx context.Context, key ContextKey, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

// HasValue checks whether the context contains a non-nil value for the given key.
func HasValue(ctx context.Context, key ContextKey) bool {
	return ctx.Value(key) != nil
}
