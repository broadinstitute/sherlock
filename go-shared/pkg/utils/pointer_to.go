package utils

// PointerTo returns a pointer to whatever you give it, so you don't need to
// define a bunch of temporary variables in tests. The fun generics make it
// always agree with the type system.
func PointerTo[T any](val T) *T {
	return &val
}
