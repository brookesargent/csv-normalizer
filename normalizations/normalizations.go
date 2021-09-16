package normalizations

import (
	"fmt"
	"strings"
	"time"
)

const layout = "1/2/06 3:04:05 PM"
const EasternTime = "America/New_York"
const PacificTime = "America/Los_Angeles"

// FormatTimestamp takes a timestamp (Pacific time assumed) and converts it to Eastern time and formats it in RFC3339
func FormatTimestamp(timestamp string) (string, error) {
	// load locations for timezone conversion
	ptLoc, err := time.LoadLocation(PacificTime)
	if err != nil {
		return "", fmt.Errorf("could not load location: %s", PacificTime)
	}
	etLoc, err := time.LoadLocation(EasternTime)
	if err != nil {
		return "", fmt.Errorf("could not load location: %s", EasternTime)
	}

	// convert to Go time
	t, err := time.ParseInLocation(layout, timestamp, ptLoc)
	if err != nil {
		return "", fmt.Errorf("could not parse timestamp: %s", timestamp)
	}

	// return timestamp string converted to ET and RFC3339 formatted
	return t.In(etLoc).Format(time.RFC3339), nil
}

// FormatZipCode pads a zip code of less than 5 characters with leading zeroes
func FormatZipCode(zip string) string {
	for len(zip) < 5 {
		zip = fmt.Sprintf("%s%s", "0", zip)
	}
	return zip
}

// FormatFullName returns the input string in uppercase
func FormatFullName(name string) string {
	return strings.ToUpper(name)
}

// DurationToSeconds takes a duration in HH:MM:SS.MS format and returns the total seconds of the duration
func DurationToSeconds(duration string) float64 {
	splitDuration := strings.Split(duration, ":")
	for i, v := range splitDuration {
		switch i {
		case 0:
			splitDuration[i] = fmt.Sprintf("%sh", v)
		case 1:
			splitDuration[i] = fmt.Sprintf("%sm", v)
		case 2:
			splitDuration[i] = fmt.Sprintf("%sms", strings.Replace(v, ".", "s", 1))
		}
	}

	duration = strings.Join(splitDuration, "")
	parsedDuration, _ := time.ParseDuration(duration)

	return parsedDuration.Seconds()
}
