package pointer

// To returns a pointer to the provided value.
func To[T any](x T) *T {
	return &x
}
