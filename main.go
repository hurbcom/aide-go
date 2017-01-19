package main

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
	HTTP_STATUS_UNPROCESSABLE_ENTITY = 422
)

const (
	// 2006 = ano com quatro digitos
	//   01 = mes com dois digitos
	//   02 = dia com dois digitos
	DATE_PATTERN_YYYYMMDD = "2006-01-02"

	// 2006 = ano com quatro digitos
	//   01 = mes com dois digitos
	//   02 = dia com dois digitos
	//   15 = hora com dois digitos (24h)
	//   04 = minuto com dois digitos
	//   05 = segundo com dois digitos
	DATE_PATTERN_YYYYMMDD_HHMMSS = "2006-01-02 15:04:05"

	// 2006 = ano com quatro digitos
	//   01 = mes com dois digitos
	//   02 = dia com dois digitos
	//   15 = hora com dois digitos (24h)
	//   04 = minuto com dois digitos
	//   05 = segundo com dois digitos
	DATE_PATTERN_YYYYMMDDTHHMMSS = "2006-01-02T15:04:05"

	// 2006 = ano com quatro digitos
	//   01 = mes com dois digitos
	//   02 = dia com dois digitos
	//   15 = hora com dois digitos (24h)
	//   04 = minuto com dois digitos
	//   05 = segundo com dois digitos
	//   Z  = UTC
	DATE_PATTERN_YYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"
)

func ToStringSlice(intslice []int) (stringSlice []string) {
	for _, i := range intslice {
		stringSlice = append(stringSlice, strconv.FormatInt(int64(i), 10))
	}
	return stringSlice
}

func ToStringSlice64(int64Slice []int64) (stringSlice []string) {
	for _, i := range int64Slice {
		stringSlice = append(stringSlice, strconv.FormatInt(i, 10))
	}
	return stringSlice
}

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

func ParseStringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.Atoi(s)
}

func ParseStringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.ParseInt(s, 10, 0)
}

func ParseDateYearMonthDay(dateString string) (time.Time, error) {
	return time.Parse(DATE_PATTERN_YYYYMMDD, dateString)
}

func DiffDays(date1 time.Time, date2 time.Time) (int64, error) {
	if !date1.IsZero() && !date2.IsZero() {
		duration := date2.Sub(date1)
		return int64(duration.Hours() / 24), nil
	}
	return 0, errors.New("invalid-dates")
}

func ParseDateStringToTime(dateString string) (*time.Time, error) {
	var result time.Time
	var err error

	if len(dateString) == 0 {
		return nil, nil
	}

	if regexp.MustCompile(`^0{4}-0{2}-0{2}[T\s]?(0{2}:0{2}:0{2})?Z?$`).MatchString(dateString) {
		fmt.Printf("ParseDateStringToTime: receiving date string zero filled. let %s as %s", dateString, result)
	} else if regexp.MustCompile(`^\d{4}\-\d{2}\-\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(DATE_PATTERN_YYYYMMDD, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(DATE_PATTERN_YYYYMMDD_HHMMSS, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`).MatchString(dateString) {
		result, err = time.Parse(DATE_PATTERN_YYYYMMDDTHHMMSS, dateString)
	} else if regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`).MatchString(dateString) {
		result, err = time.Parse(DATE_PATTERN_YYYYMMDDTHHMMSSZ, dateString)
	} else {
		err = errors.New(fmt.Sprintf("invalid date format - %+v", dateString))
	}

	return &result, err
}

func RemoveNanoseconds(date time.Time) (time.Time, error) {
	dateWithoutNSecs, err := ParseDateStringToTime(date.Format(time.RFC3339))
	if err != nil {
		return date, err
	}
	return *dateWithoutNSecs, nil
}

func ParseIntToBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func ParseStringToBool(s string) bool {
	if s == "1" {
		return true
	}
	return false
}

func ParseBoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func CheckStringJsonData(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

func CheckInt64JsonData(i int64) *int64 {
	if i > 0 {
		return &i
	}
	return nil
}

func CheckFloat64JsonData(f float64) *float64 {
	if f > 0 {
		return &f
	}
	return nil
}

func GetByteArrayAndBufferFromRequestBody(body io.ReadCloser) ([]byte, *bytes.Buffer, error) {
	defer body.Close()
	byteArray, err := ioutil.ReadAll(body)
	if err != nil {
		return []byte{}, nil, err
	}
	buffer := bytes.NewBuffer(byteArray)
	return byteArray, buffer, nil
}

func GetOnlyNumbers(s *string) *string {
	return GetOnlyNumbersOrSpecial(s, "")
}

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
		fmt.Printf("Printing request Body: %+v", groups[0])
		return &groups[0]
	}

	fmt.Printf("No body to print on request Body")
	return nil
}

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
		fmt.Printf("Printing response Body: %+v", groups[0])
		return &groups[0]
	}

	fmt.Printf("No body to print on response Body")
	return nil
}

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

func IsArray(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Slice
}

func Join(sep string, args ...interface{}) string {
	var buf bytes.Buffer
	var str Stringer
	var ok bool
	var elements []interface{}

	for _, arg := range args {
		if IsArray(arg) {
			valueArg := reflect.ValueOf(arg)
			for j := 0; j < valueArg.Len(); j++ {
				elements = append(elements, valueArg.Index(j).Interface())
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

func BeginningOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
