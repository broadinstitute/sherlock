package role_propagation

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"reflect"
)

// shouldPropagate determines whether this propagator should act on the given role.
//
// It uses a tiny bit of reflection: if the grant is nil or the zero value of its type, we will not propagate.
// This is because we often explicitly store Go zero values in the database to represent "unset", because that's
// way easier for us to work with than null itself -- Gorm will ignore nulls by default to try to fit with Go's
// zero value semantics.
func (p *propagatorImpl[Grant, Identifier, Fields]) shouldPropagate(role models.Role) (shouldPropagate bool, grant Grant) {
	grantPointer := p.getGrant(role)
	if grantPointer == nil || reflect.ValueOf(*grantPointer).IsZero() {
		return false, grant
	} else {
		return true, *grantPointer
	}
}
