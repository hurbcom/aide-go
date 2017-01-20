package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	httpStatusUnprocessableEntity = 422

	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	datePatternYYYYMMDD = "2006-01-02"

	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	datePatternYYYYMMDDHHMMSS = "2006-01-02 15:04:05"

	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	datePatternYYYYMMDDTHHMMSS = "2006-01-02T15:04:05"

	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	//   Z  = UTC
	datePatternYYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"
)

// ToStringSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ToStringSlice(intslice []int) (stringSlice []string) {
	for _, i := range intslice {
		stringSlice = append(stringSlice, strconv.FormatInt(int64(i), 10))
	}
	return stringSlice
}

// ToStringSlice64 REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ToStringSlice64(int64Slice []int64) (stringSlice []string) {
	for _, i := range int64Slice {
		stringSlice = append(stringSlice, strconv.FormatInt(i, 10))
	}
	return stringSlice
}

// ToInt64Slice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ToInt64Slice(stringSlice []string) (int64Slice []int64) {
	for _, i := range stringSlice {
		intI, err := ParseStringToInt64(i)
		if err != nil {
			continue
		}
		int64Slice = append(int64Slice, intI)
	}
	return int64Slice
}

// ParseStringToInt REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseStringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.Atoi(s)
}

// ParseStringToInt64 REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseStringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.ParseInt(s, 10, 0)
}

// ParseDateYearMonthDay REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseDateYearMonthDay(dateString string) (time.Time, error) {
	return time.Parse(datePatternYYYYMMDD, dateString)
}

// DiffDays REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DiffDays(date1 time.Time, date2 time.Time) (int64, error) {
	if !date1.IsZero() && !date2.IsZero() {
		duration := date2.Sub(date1)
		return int64(duration.Hours() / 24), nil
	}
	return 0, errors.New("invalid-dates")
}

// ParseDateStringToTime REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseDateStringToTime(dateString string) (*time.Time, error) {
	var result time.Time
	var err error

	if len(dateString) == 0 {
		return nil, nil
	}

	if regexp.MustCompile(`^0{4}-0{2}-0{2}[T\s]?(0{2}:0{2}:0{2})?Z?$`).MatchString(dateString) {
		fmt.Printf("ParseDateStringToTime: receiving date string zero filled. let %s as %s", dateString, result)
	} else if regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(datePatternYYYYMMDD, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(datePatternYYYYMMDDHHMMSS, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(datePatternYYYYMMDDTHHMMSS, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`).MatchString(dateString) {
		result, err = time.Parse(datePatternYYYYMMDDTHHMMSSZ, dateString)
	} else {
		err = errors.New(fmt.Sprintf("ParseDateStringToTime: invalid date format - %+v", dateString))
	}

	return &result, err
}

// RemoveNanoseconds REQUIRE THEM TO DOCUMENT THIS FUNCTION
func RemoveNanoseconds(date time.Time) (time.Time, error) {
	dateWithoutNSecs, err := ParseDateStringToTime(date.Format(time.RFC3339))
	if err != nil {
		return date, err
	}
	return *dateWithoutNSecs, nil
}

// ParseIntToBool REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseIntToBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

// ParseStringToBool REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseStringToBool(s string) bool {
	if s == "1" {
		return true
	}
	return false
}

// ParseBoolToString REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseBoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// CheckStringJsonData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckStringJsonData(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

// CheckInt64JsonData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckInt64JsonData(i int64) *int64 {
	if i > 0 {
		return &i
	}
	return nil
}

// CheckFloat64JsonData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckFloat64JsonData(f float64) *float64 {
	if f > 0 {
		return &f
	}
	return nil
}

// GetByteArrayAndBufferFromRequestBody REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetByteArrayAndBufferFromRequestBody(body io.ReadCloser) ([]byte, *bytes.Buffer, error) {
	defer body.Close()
	byteArray, err := ioutil.ReadAll(body)
	if err != nil {
		return []byte{}, nil, err
	}
	buffer := bytes.NewBuffer(byteArray)
	return byteArray, buffer, nil
}

// GetOnlyNumbers REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetOnlyNumbers(s *string) *string {
	return GetOnlyNumbersOrSpecial(s, "")
}

// GetOnlyNumbersOrSpecial REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetOnlyNumbersOrSpecial(s *string, sp string) *string {
	if s == nil {
		return s
	} else {
		specials := ""
		if len(sp) > 0 {
			for _, item := range strings.Split(sp, "") {
				specials = specials + `\` + item
			}
		}
		pattern := fmt.Sprintf(`[^%s0-9]`, specials)
		r := regexp.MustCompile(pattern)
		result := r.ReplaceAllString(*s, "")
		return &result
	}
}

// GetStringBodyHttpRequest REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHttpRequest(r *http.Request) *string {
	if r == nil {
		return nil
	}

	headers, _ := httputil.DumpRequest(r, false)
	headersAndBody, _ := httputil.DumpRequest(r, true)
	body := headersAndBody[len(headers):]
	string_body := string(body)

	re := regexp.MustCompile(`(?s)(.*)`)
	groups := re.FindStringSubmatch(string_body)

	if len(groups) > 0 {
		fmt.Printf("GetStringBodyHttpRequest: printing request Body: %+v", groups[0])
		return &groups[0]
	}

	fmt.Printf("GetStringBodyHttpRequest: no body to print on request Body")
	return nil
}

// GetStringBodyHttpRequestJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHttpRequestJSON(r *http.Request) *string {
	result := GetStringBodyHttpRequest(r)
	if result != nil {
		re := regexp.MustCompile(`({.*})`)
		groups := re.FindStringSubmatch(*result)
		if len(groups) > 0 {
			return &groups[0]
		}
		return result
	}
	return nil
}

// GetStringBodyHttpResponse REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHttpResponse(r *http.Response) *string {
	if r == nil {
		return nil
	}

	headers, _ := httputil.DumpResponse(r, false)
	headersAndBody, _ := httputil.DumpResponse(r, true)
	body := headersAndBody[len(headers):]
	string_body := string(body)

	re := regexp.MustCompile(`(?s)(.*)`)
	groups := re.FindStringSubmatch(string_body)

	if len(groups) > 0 {
		fmt.Printf("GetStringBodyHttpResponse: printing response Body: %+v", groups[0])
		return &groups[0]
	}

	fmt.Printf("GetStringBodyHttpResponse: no body to print on response Body")
	return nil
}

// GetStringBodyHttpResponseJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHttpResponseJSON(r *http.Response) *string {
	result := GetStringBodyHttpResponse(r)
	if result != nil {
		re := regexp.MustCompile(`({.*})`)
		groups := re.FindStringSubmatch(*result)
		if len(groups) > 0 {
			return &groups[0]
		}
		return result
	}
	return nil
}

// ParseIntOrReturnZero REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseIntOrReturnZero(s string) int {
	integer, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(integer)
}

type Stringer interface {
	String() string
}

// IsArray REQUIRE THEM TO DOCUMENT THIS FUNCTION
func IsArray(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Slice
}

// IsString REQUIRE THEM TO DOCUMENT THIS FUNCTION
func IsString(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.String
}

// IsPointer REQUIRE THEM TO DOCUMENT THIS FUNCTION
func IsPointer(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Ptr
}

// Join REQUIRE THEM TO DOCUMENT THIS FUNCTION
func Join(sep string, args ...interface{}) string {
	var buf bytes.Buffer
	var str Stringer
	var ok bool
	var elements []interface{}

	for _, arg := range args {
		if arg == nil {
			continue
		}

		if IsArray(arg) {
			valueArg := reflect.ValueOf(arg)
			for j := 0; j < valueArg.Len(); j++ {
				elements = append(elements, valueArg.Index(j).Interface())
			}
		} else if IsString(arg) {
			if len(arg.(string)) > 0 {
				elements = append(elements, arg)
			}
		} else if IsPointer(arg) {
			valueArg := reflect.ValueOf(arg)
			if valueArg.Elem().IsValid() {
				elements = append(elements, valueArg.Elem())
			}
		} else {
			elements = append(elements, arg)
		}
	}

	for i, arg := range elements {
		if i > 0 {
			buf.WriteString(sep)
		}

		if str, ok = arg.(Stringer); ok {
			buf.WriteString(str.String())
		} else {
			fmt.Fprint(&buf, arg)
		}
	}

	return buf.String()
}

// BeginningOfToday REQUIRE THEM TO DOCUMENT THIS FUNCTION
func BeginningOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
