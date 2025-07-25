package role_propagation

import (
	"reflect"

	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

func (p *propagatorImpl[Grant, Identifier, Fields]) getAndFilterGrants(role models.Role) []Grant {
	filteredGrants := make([]Grant, 0)
	if unfilteredGrants := p.getGrants(role); unfilteredGrants != nil {
		for _, grantPointer := range unfilteredGrants {
			if grantPointer != nil && !reflect.ValueOf(*grantPointer).IsZero() {
				filteredGrants = append(filteredGrants, *grantPointer)
			}
		}
	}
	return filteredGrants
}
