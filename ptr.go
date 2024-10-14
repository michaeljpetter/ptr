// Package ptr implements a set of utilities
// for working with pointers in Go.
package ptr

// To returns a typed pointer to a copy of the given value.
func To[T any](t T) *T {
	return &t
}

// OrZero returns the given pointer if non-nil, otherwise
// a pointer to the zero value of the referenced type.
func OrZero[T any](p *T) *T {
	if p == nil {
		var zero T
		return &zero
	}

	return p
}
