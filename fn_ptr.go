package glodash

// Ptr returns a pointer to the provided value.
// This is useful when you need a pointer to a constant
// or literal value.
func Ptr[T any](value T) *T {
	return &value
}
