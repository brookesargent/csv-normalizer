package normalizations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatZipCode(t *testing.T) {
	expected := "01125"
	actual := FormatZipCode("1125")
	assert.Equal(t, expected, actual)

	expected = "00125"
	actual = FormatZipCode("125")
	assert.Equal(t, expected, actual)

	expected = "11125"
	actual = FormatZipCode("11125")
	assert.Equal(t, expected, actual)
}

func TestFormatFullName(t *testing.T) {
	expected := "BROOKE SARGENT"
	actual := FormatFullName("brooke sargent")
	assert.Equal(t, expected, actual)
}

func TestDurationToSeconds(t *testing.T) {
	expected := 5012.123
	actual := DurationToSeconds("1:23:32.123")
	assert.Equal(t, expected, actual)
}
