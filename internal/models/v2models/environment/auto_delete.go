package environment

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type AutoDelete struct {
	Enabled bool      `json:"enabled" form:"enabled" default:"false"`
	After   time.Time `json:"after" form:"after"`
}

// Validate returns an error if the given auto-delete configuration is not valid
func (a AutoDelete) Validate() error {
	isEnabled := a.Enabled
	nonZeroTime := !a.After.IsZero()

	if isEnabled != nonZeroTime {
		return fmt.Errorf("invalid auto-delete configuration: %#v", a)
	}
	return nil
}

// Value implement the driver.Valuer interface (for serializing AutoDelete to database)
func (a AutoDelete) Value() (driver.Value, error) {
	if !a.Enabled {
		// a"disabled" autodelete is stored as null in database
		return nil, nil
	}
	// non-null "delete-after" timestamp indicates auto-delete is enabled
	return a.After, nil
}

// Scan implement the sql.Scanner interface (for deserializing AutoDelete from the database)
func (a *AutoDelete) Scan(value interface{}) error {
	// sql package already has a "possibly-null" time type that does what we need
	var maybeNull sql.NullTime
	if err := maybeNull.Scan(value); err != nil {
		return err
	}
	a.Enabled = maybeNull.Valid
	a.After = maybeNull.Time
	return nil
}
