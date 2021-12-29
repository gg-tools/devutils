package timestamp

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	dateFormat = "2006-01-02"
	dateLen    = len(dateFormat)
	timeFormat = "2006-01-02 15:04:05"
	timeLen    = len(timeFormat)
)

func parseTimestamp(str string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestamp, 0), nil
}

func parseTime(arg string) (t time.Time, err error) {
	switch len(arg) {
	case dateLen:
		return time.ParseInLocation(dateFormat, arg, time.Local)
	case timeLen:
		return time.ParseInLocation(timeFormat, arg, time.Local)
	default:
		return time.Time{}, errors.New("bad timestamp format")
	}

}

func str2time(s ...string) (t time.Time, err error) {
	exp := regexp.MustCompile(`^(\+|-)?(\d+)([ymdhis])$`)

	t = time.Now()
	for i, ss := range s {
		parts := exp.FindStringSubmatch(ss)
		if len(parts) != 4 {
			err = fmt.Errorf("bad arg at %d: %s", i, ss)
			return
		}
		num, _ := strconv.Atoi(parts[1] + parts[2])
		measure := parts[3]

		switch measure {
		case "y":
			t = t.AddDate(num, 0, 0)
		case "m":
			t = t.AddDate(0, num, 0)
		case "d":
			t = t.AddDate(0, 0, num)
		case "h":
			t = t.Add(time.Duration(num) * time.Hour)
		case "i":
			t = t.Add(time.Duration(num) * time.Minute)
		case "s":
			t = t.Add(time.Duration(num) * time.Second)
		}
	}

	return t, nil
}

func printTimestamp(t time.Time) {
	fmt.Printf("%d (%s)\n", t.Unix(), t.Format(timeFormat))
}

func printTime(t time.Time) {
	fmt.Printf("%s (%d)\n", t.Format(timeFormat), t.Unix())
}
