package normalizations

import (
	"fmt"
	"strings"
	"time"
)

const layout = "1/2/06 3:04:05 PM"

func FormatTimestamp(timestamp string) string {
	// load locations for timezone conversion
	ptLoc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Could not parse location")
	}
	etLoc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Could not parse location")
	}

	// convert to Go time
	t, err := time.ParseInLocation(layout, timestamp, ptLoc)
	if err != nil {
		fmt.Println(err)
	}

	// return timestamp string converted to ET and RFC3339 formatted
	return t.In(etLoc).Format(time.RFC3339)
}

func FormatZipCode(zip string) string {
	for len(zip) < 5 {
		zip = fmt.Sprintf("%s%s", "0", zip)
	}
	return zip
}

func FormatFullName(name string) string {
	return strings.ToUpper(name)
}

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
