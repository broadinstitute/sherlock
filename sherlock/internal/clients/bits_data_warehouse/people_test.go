package bits_data_warehouse

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_calculatePersonColumns(t *testing.T) {
	// Check that reflection doesn't panic
	require.NotPanics(t, func() { calculatePersonColumns() })

	// Check that we have "group" properly escaped because it causes a confusing error
	// if it isn't (it says expected "BY" but got ",", which makes sense if you know
	// that "group" is both a column here and a reserved word in SQL).
	assert.Contains(t, personColumns, "`group`")
}
