package aidego

import (
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	// HTTPStatusUnprocessableEntity Deprecated: use http.StatusUnprocessableEntity instead
	HTTPStatusUnprocessableEntity = 422

	// DatePatternYYYYMMDD
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	DatePatternYYYYMMDD = "2006-01-02"

	// DatePatternYYYYMMDDHHMMSS
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	DatePatternYYYYMMDDHHMMSS = "2006-01-02 15:04:05"

	// DatePatternYYYYMMDDTHHMMSS
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	DatePatternYYYYMMDDTHHMMSS = "2006-01-02T15:04:05"

	// DatePatternYYYYMMDDTHHMMSS
	// ISO 8601 format with timezone offset
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	//   07 = Timezone offset hours with two digits
	//   00 = Timezone offset minutes with two digits
	DatePatternYYYYMMDDTHHMMSSOffset = "2006-01-02T15:04:05-07:00"

	DatePatternYYYYMMDDTHHMMSSZ = time.RFC3339 // NOTE: backward compatibility
)

var (
	sRFC3339NanoWithoutTAndTimezone = strings.ReplaceAll(
		strings.ReplaceAll(time.RFC3339Nano, `T`, ` `),
		`Z07:00`, ``)

	regexpDatePatternYYYYMMDD = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}$`)

	regexpDatePatternYYYYMMDDHHMMSS = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)

	regexpDatePatternYYYYMMDDTHHMMSS = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`)

	regexpDatePatternYYYYMMDDTHHMMSSOffset *regexp.Regexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$([\+|\-]([01]\d|2[0-3]):[0-5]\d)`)

	sTimezone = `(([Zz])|([\+|\-]([01]\d|2[0-3]):[0-5]\d))`

	sRegexpRFC3339 = `^(\d+)-(0[1-9]|1[012])-(0[1-9]|[12]\d|3[01])` +
		`[Tt]([01]\d|2[0-3]):([0-5]\d):([0-5]\d|60)(\.\d+)?`

	regexpRFC3339 = regexp.MustCompile(sRegexpRFC3339 + sTimezone + `$`)

	sRegexpRFC3339NanoWithoutTAndTimezone = strings.ReplaceAll(sRegexpRFC3339, `[Tt]`, `\s`) + `\.\d{1,9}` + sTimezone + `?$`

	regexpRFC3339NanoWithoutTAndTimezone = regexp.MustCompile(sRegexpRFC3339NanoWithoutTAndTimezone)
)

// ParseDateYearMonthDay
func ParseDateYearMonthDay(dateString string) (time.Time, error) {
	return time.Parse(DatePatternYYYYMMDD, dateString)
}

// DiffDays
func DiffDays(date1 time.Time, date2 time.Time) (int64, error) {
	if !date1.IsZero() && !date2.IsZero() {
		duration := date2.Sub(date1)
		days := math.Ceil(duration.Hours() / 24)
		return int64(days), nil
	}
	return 0, errors.Errorf("invalid-dates: %v or %v is invalid", date1, date2)
}

// ParseDateStringToTime
func ParseDateStringToTime(dateString string) (*time.Time, error) {
	if len(dateString) == 0 {
		return nil, errors.New("empty date format")
	}

	matchers := map[string]*regexp.Regexp{
		DatePatternYYYYMMDD:              regexpDatePatternYYYYMMDD,
		DatePatternYYYYMMDDHHMMSS:        regexpDatePatternYYYYMMDDHHMMSS,
		DatePatternYYYYMMDDTHHMMSS:       regexpDatePatternYYYYMMDDTHHMMSS,
		DatePatternYYYYMMDDTHHMMSSZ:      regexpRFC3339,
		DatePatternYYYYMMDDTHHMMSSOffset: regexpDatePatternYYYYMMDDTHHMMSSOffset,
		sRFC3339NanoWithoutTAndTimezone:  regexpRFC3339NanoWithoutTAndTimezone,
	}

	for k, v := range matchers {
		if v.MatchString(dateString) {
			result, err := time.Parse(k, dateString)
			if err != nil {
				return nil, errors.Errorf("using pattern %s result error: %v", k, err)
			}
			return &result, nil
		}
	}

	return nil, errors.Errorf("invalid date format - %+v", dateString)
}

// ParseDateStringToTimeIn parses a date string into time type in a specific location
func ParseDateStringToTimeIn(dateString string, loc *time.Location) (*time.Time, error) {
	if len(dateString) == 0 {
		return nil, errors.New("empty date format")
	}

	matchers := map[string]*regexp.Regexp{
		DatePatternYYYYMMDD:              regexpDatePatternYYYYMMDD,
		DatePatternYYYYMMDDHHMMSS:        regexpDatePatternYYYYMMDDHHMMSS,
		DatePatternYYYYMMDDTHHMMSS:       regexpDatePatternYYYYMMDDTHHMMSS,
		DatePatternYYYYMMDDTHHMMSSZ:      regexpRFC3339,
		DatePatternYYYYMMDDTHHMMSSOffset: regexpDatePatternYYYYMMDDTHHMMSSOffset,
		sRFC3339NanoWithoutTAndTimezone:  regexpRFC3339NanoWithoutTAndTimezone,
	}

	for k, v := range matchers {
		if v.MatchString(dateString) {
			result, err := time.ParseInLocation(k, dateString, loc)
			if err != nil {
				return nil, errors.Errorf("using pattern %s result error: %v", k, err)
			}
			return &result, nil
		}
	}

	return nil, errors.Errorf("invalid date format - %+v", dateString)
}

// RemoveNanoseconds
func RemoveNanoseconds(date time.Time) (time.Time, error) {
	dateWithoutNSecs, err := ParseDateStringToTime(date.Format(time.RFC3339))
	if err != nil {
		return date, err
	}
	return *dateWithoutNSecs, nil
}

// BeginningOfToday
func BeginningOfToday() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// BeginningOfTodayIn
func BeginningOfTodayIn(loc *time.Location) time.Time {
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}