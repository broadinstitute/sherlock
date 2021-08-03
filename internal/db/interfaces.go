package db

// Selector exposes an interface for abstracting Selection operations from
// some datastore.
type Selector interface {
	Select(dest interface{}, query string, args ...interface{}) error
}
