package ctxext

import "context"

// CxUser retrieves the user name from the context
func CxUser(ctx context.Context) string {
	return GetValue[string](ctx, UserName)
}

// CxSession retrieves the session from the context
func CxSession(ctx context.Context) string {
	return GetValue[string](ctx, SessionID)
}

// CxDuration retrieves the cache duration from the context
func CxDuration(ctx context.Context) int {
	return GetValue[int](ctx, CacheDuration)
}

// CxSecret retrieves the secret keys from the context
func CxSecret(ctx context.Context) string {
	return GetValue[string](ctx, SecretKeys)
}

// CxMax retrieves the max records from the context
func CxMax(ctx context.Context) int {
	return GetValue[int](ctx, MaxRecords)
}

// CxSize retrieves the page size from the context
func CxSize(ctx context.Context) int {
	return GetValue[int](ctx, PageSize)
}

// CxUserPtr retrieves the user name from the context and returns a pointer
func CxUserPtr(ctx context.Context) *string {
	val := GetValue[string](ctx, UserName)
	return &val
}

// CxSessionPtr retrieves the session from the context and returns a pointer
func CxSessionPtr(ctx context.Context) *string {
	val := GetValue[string](ctx, SessionID)
	return &val
}

// CxDurationPtr retrieves the cache duration from the context and returns a pointer
func CxDurationPtr(ctx context.Context) *int {
	val := GetValue[int](ctx, CacheDuration)
	return &val
}

// CxSecretPtr retrieves the secret keys from the context and returns a pointer
func CxSecretPtr(ctx context.Context) *string {
	val := GetValue[string](ctx, SecretKeys)
	return &val
}

// CxMaxPtr retrieves the max records from the context and returns a pointer
func CxMaxPtr(ctx context.Context) *int {
	val := GetValue[int](ctx, MaxRecords)
	return &val
}

// CxSizePtr retrieves the page size from the context and returns a pointer
func CxSizePtr(ctx context.Context) *int {
	val := GetValue[int](ctx, PageSize)
	return &val
}
