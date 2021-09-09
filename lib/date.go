package lib

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

	DatePatternYYYYMMDDTHHMMSSZ = time.RFC3339 // NOTE: backward compatibility
)

var (
	RFC3339NanoWithoutTAndTimezone string = strings.ReplaceAll(
		strings.ReplaceAll(time.RFC3339Nano, `T`, ` `),
		`Z07:00`, ``)

	regexpDatePatternYYYYMMDD *regexp.Regexp = regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}$`)

	regexpDatePatternYYYYMMDDHHMMSS *regexp.Regexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)

	regexpDatePatternYYYYMMDDTHHMMSS *regexp.Regexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`)

	timezone string = `(([Zz])|([\+|\-]([01]\d|2[0-3]):[0-5]\d))`

	sRegexpRFC3339 string = `^(\d+)-(0[1-9]|1[012])-(0[1-9]|[12]\d|3[01])` +
		`[Tt]([01]\d|2[0-3]):([0-5]\d):([0-5]\d|60)(\.\d+)?`

	regexpRFC3339 *regexp.Regexp = regexp.MustCompile(sRegexpRFC3339 + timezone + `$`)

	sRegexpRFC3339NanoWithoutTAndTimezone = strings.ReplaceAll(sRegexpRFC3339, `[Tt]`, `\s`) + `\.\d{1,9}` + timezone + `?$`

	regexpRFC3339NanoWithoutTAndTimezone *regexp.Regexp = regexp.MustCompile(sRegexpRFC3339NanoWithoutTAndTimezone)

	regexpCommaAlphaNum *regexp.Regexp = regexp.MustCompile(`[^A-Za-z0-9,]`)
)

// ParseDateYearMonthDay REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseDateYearMonthDay(dateString string) (time.Time, error) {
	return time.Parse(DatePatternYYYYMMDD, dateString)
}

// DiffDays REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DiffDays(date1 time.Time, date2 time.Time) (int64, error) {
	if !date1.IsZero() && !date2.IsZero() {
		duration := date2.Sub(date1)
		days := math.Ceil(duration.Hours() / 24)
		return int64(days), nil
	}
	return 0, errors.Errorf("invalid-dates: %v or %v is invalid", date1, date2)
}

// ParseDateStringToTime REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseDateStringToTime(dateString string) (*time.Time, error) {
	if len(dateString) == 0 {
		return nil, errors.New("empty date format")
	}

	matchers := map[string]*regexp.Regexp{
		DatePatternYYYYMMDD:                    regexpDatePatternYYYYMMDD,
		DatePatternYYYYMMDDHHMMSS:              regexpDatePatternYYYYMMDDHHMMSS,
		DatePatternYYYYMMDDTHHMMSS:             regexpDatePatternYYYYMMDDTHHMMSS,
		string(time.RFC3339):                   regexpRFC3339,
		string(RFC3339NanoWithoutTAndTimezone): regexpRFC3339NanoWithoutTAndTimezone,
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

// RemoveNanoseconds REQUIRE THEM TO DOCUMENT THIS FUNCTION
func RemoveNanoseconds(date time.Time) (time.Time, error) {
	dateWithoutNSecs, err := ParseDateStringToTime(date.Format(time.RFC3339))
	if err != nil {
		return date, err
	}
	return *dateWithoutNSecs, nil
}

// BeginningOfToday REQUIRE THEM TO DOCUMENT THIS FUNCTION
func BeginningOfToday() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// BeginningOfTodayIn REQUIRE THEM TO DOCUMENT THIS FUNCTION
func BeginningOfTodayIn(loc *time.Location) time.Time {
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}
