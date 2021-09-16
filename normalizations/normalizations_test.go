// +build unit

package normalizations

import (
	"fmt"
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

func TestFormatTimestampSuccess(t *testing.T) {
	expected := "2016-03-13T03:01:00-04:00"
	actual, err := FormatTimestamp("3/12/16 11:01:00 PM")
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func TestFormatTimestampError(t *testing.T) {
	ts := "3/12/16 11:01 PM"
	result, err := FormatTimestamp(ts)
	fmt.Println(err)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("could not parse timestamp: %s", ts), err.Error())
	assert.Equal(t, "", result)
}
