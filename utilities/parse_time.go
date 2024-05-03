package utilities

import "time"

func ParseTime(date string) (time.Time, error) {
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
