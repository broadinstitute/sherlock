package resource_prefix

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"strings"
	"testing"
)

func TestGenerateResourcePrefix(t *testing.T) {
	sb := strings.Builder{}
	sb.Grow(4)
	tests := []struct {
		input  uint64
		output string
	}{
		{0, "aaaa"},
		{1, "aaab"},
		{2, "aaac"},
		{possibleCombinations - 2, "z998"},
		{possibleCombinations - 1, "z999"},
		{possibleCombinations, "aaaa"},
		{possibleCombinations + 1, "aaab"},
		{possibleCombinations + 2, "aaac"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d to %s", tt.input, tt.output), func(t *testing.T) {
			GenerateResourcePrefix(&sb, tt.input)
			testutils.AssertNoDiff(t, tt.output, sb.String())
			sb.Reset()
		})
	}
}
