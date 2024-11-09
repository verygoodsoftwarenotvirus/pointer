package pointer

// To returns a pointer to a value.
func To[T any](x T) *T {
	return &x
}

// Deref returns the value of a pointer.
func Deref[T any](x *T) T {
	if x == nil {
		var zero T
		return zero
	}
	return *x
}

// ForSlice returns the pointer for every element in a slice.
func ForSlice[T any](x []T) []*T {
	if x == nil {
		return []*T{}
	}

	y := make([]*T, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	return y
}

// DerefSlice returns the value of a pointer for every element in a slice.
func DerefSlice[T any](x []*T) []T {
	if x == nil {
		return []T{}
	}

	y := make([]T, len(x))
	for i := range x {
		y[i] = *x[i]
	}
	return y
}

// ForMap returns the pointer for every element in a map.
func ForMap[K comparable, V any](x map[K]V) map[K]*V {
	if x == nil {
		return map[K]*V{}
	}

	y := make(map[K]*V, len(x))
	for k, v := range x {
		y[k] = &v
	}
	return y
}

// DerefMap returns the value of a pointer for every element in a map.
func DerefMap[K comparable, V any](x map[K]*V) map[K]V {
	if x == nil {
		return map[K]V{}
	}

	y := make(map[K]V, len(x))
	for k, v := range x {
		y[k] = *v
	}
	return y
}
