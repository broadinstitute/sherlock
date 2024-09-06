package role_propagation

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"strings"
)

// a convenience function helpful for handling grant fields on roles that we treat as comma-separated lists.
// This function doesn't need to handle filtering the list at all because that's generic behavior that
// propagatorImpl.getAndFilterGrants will handle.
func splitStringPointerOnCommas(s *string) []*string {
	if s == nil {
		return nil
	} else {
		return utils.Map(strings.Split(*s, ","), func(s string) *string { return utils.PointerTo(strings.TrimSpace(s)) })
	}
}
