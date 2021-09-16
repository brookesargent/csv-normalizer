package normalizations

import (
	"fmt"
	"strings"
	"time"
)

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
