package db

import "database/sql"

// Selector exposes an interface for abstracting Selection operations against
// some datastore. Primary purpose is to support mocking of the postgres backend in
// unit tests
type Selector interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

// Executor is an interface representing the capability to execute a query
// against some backend  data storage.
type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Preparer represents the capabilty to create a prepared query against some
// backend data store. The primary purpose of this interface is to support the
// mocking of a postgres data store in unit tests
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// Closer represents the capability to close a connection to a backend datastore.
// Primary purpose is to support mocking a postgres db in unit tests
type Closer interface {
	Close() error
}

// Querier is a composite interface consolidating all the methods
// sherlock requires from it's backend datastore into a single type.
// Name subject to change.
type Querier interface {
	Selector
	Preparer
	Executor
	Closer
}
