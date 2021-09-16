package normalizations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatZipCode(t *testing.T) {
	expected := "01125"
	actual := FormatZipCode("1125")
	assert.Equal(t, expected, actual)
}
