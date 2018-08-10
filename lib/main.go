package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
)

const (
	// HTTPStatusUnprocessableEntity REQUIRE THEM TO DOCUMENT THIS CONST
	HTTPStatusUnprocessableEntity = 422

	// DatePatternYYYYMMDD REQUIRE THEM TO DOCUMENT THIS CONST
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	DatePatternYYYYMMDD = "2006-01-02"

	// DatePatternYYYYMMDDHHMMSS REQUIRE THEM TO DOCUMENT THIS CONST
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	DatePatternYYYYMMDDHHMMSS = "2006-01-02 15:04:05"

	// DatePatternYYYYMMDDTHHMMSS REQUIRE THEM TO DOCUMENT THIS CONST
	// 2006 = Year with four digits
	//   01 = Month with two digits
	//   02 = Day with two digits
	//   15 = Hour with two digits (24h)
	//   04 = Minute with two digits
	//   05 = Seconds with two digits
	DatePatternYYYYMMDDTHHMMSS = "2006-01-02T15:04:05"

	DatePatternYYYYMMDDTHHMMSSZ = time.RFC3339
)

var (
	regexpDatePatternZeroFilled      *regexp.Regexp
	regexpDatePatternYYYYMMDD        *regexp.Regexp
	regexpDatePatternYYYYMMDDHHMMSS  *regexp.Regexp
	regexpDatePatternYYYYMMDDTHHMMSS *regexp.Regexp
	regexpRFC3339                    *regexp.Regexp
	regexpRFC3339WithTime            *regexp.Regexp
)

func init() {
	regexpDatePatternZeroFilled, _ = regexp.Compile(`^0{4}-0{2}-0{2}[T\s]?(0{2}:0{2}:0{2})?Z?$`)
	regexpDatePatternYYYYMMDD, _ = regexp.Compile(`^\d{4}\-\d{2}\-\d{2}$`)
	regexpDatePatternYYYYMMDDHHMMSS, _ = regexp.Compile(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	regexpDatePatternYYYYMMDDTHHMMSS, _ = regexp.Compile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`)
	regexpRFC3339, _ = regexp.Compile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
	regexpRFC3339WithTime, _ = regexp.Compile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\+\-]{1}\d{2}:\d{2}$`)
}

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

// ToIntSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ToIntSlice(stringSlice []string) (intSlice []int) {
	for _, i := range stringSlice {
		intI, err := ParseStringToInt(i)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, intI)
	}
	return intSlice
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
	return time.Parse(DatePatternYYYYMMDD, dateString)
}

// DiffDays REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DiffDays(date1 time.Time, date2 time.Time) (int64, error) {
	if !date1.IsZero() && !date2.IsZero() {
		duration := date2.Sub(date1)
		days := math.Ceil(duration.Hours() / 24)
		return int64(days), nil
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

	if regexpDatePatternZeroFilled.MatchString(dateString) {
		fmt.Printf("ParseDateStringToTime: receiving date string zero filled. let %s as %s", dateString, result)
	} else if regexpDatePatternYYYYMMDD.MatchString(dateString) {
		result, err = time.Parse(DatePatternYYYYMMDD, dateString)
	} else if regexpDatePatternYYYYMMDDHHMMSS.MatchString(dateString) {
		result, err = time.Parse(DatePatternYYYYMMDDHHMMSS, dateString)
	} else if regexpDatePatternYYYYMMDDTHHMMSS.MatchString(dateString) {
		result, err = time.Parse(DatePatternYYYYMMDDTHHMMSS, dateString)
	} else if regexpRFC3339.MatchString(dateString) {
		result, err = time.Parse(time.RFC3339, dateString)
	} else if regexpRFC3339WithTime.MatchString(dateString) {
		result, err = time.Parse(time.RFC3339, dateString)
	} else {
		err = fmt.Errorf("ParseDateStringToTime: invalid date format - %+v", dateString)
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

// CheckStringJSONData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckStringJSONData(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

// CheckInt64JSONData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckInt64JSONData(i int64) *int64 {
	if i > 0 {
		return &i
	}
	return nil
}

// CheckFloat64JSONData REQUIRE THEM TO DOCUMENT THIS FUNCTION
func CheckFloat64JSONData(f float64) *float64 {
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
	}
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

// GetStringBodyHTTPRequest REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPRequest(r *http.Request) *string {
	if r == nil {
		return nil
	}
	headers, _ := httputil.DumpRequest(r, false)
	headersAndBody, _ := httputil.DumpRequest(r, true)
	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPRequestJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPRequestJSON(r *http.Request) *string {
	if r == nil {
		return nil
	}
	headers, _ := httputil.DumpRequest(r, false)
	headersAndBody, _ := httputil.DumpRequest(r, true)
	body := bytes.TrimSpace(headersAndBody[len(headers):])
	if len(body) > 0 {
		start := bytes.IndexAny(body, "{")
		end := bytes.LastIndexAny(body, "}")
		r := string(body[start : end+1])
		return &r
	}
	return nil
}

// GetStringBodyHTTPResponse REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPResponse(r *http.Response) *string {
	if r == nil {
		return nil
	}
	headers, _ := httputil.DumpResponse(r, false)
	headersAndBody, _ := httputil.DumpResponse(r, true)
	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPResponseJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPResponseJSON(r *http.Response) *string {
	if r == nil {
		return nil
	}
	headers, _ := httputil.DumpResponse(r, false)
	headersAndBody, _ := httputil.DumpResponse(r, true)
	body := bytes.TrimSpace(headersAndBody[len(headers):])
	if len(body) > 0 {
		start := bytes.IndexAny(body, "{")
		end := bytes.LastIndexAny(body, "}")
		r := string(body[start : end+1])
		return &r
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

// Stringer REQUIRE THEM TO DOCUMENT THIS TYPE
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
		if str := cast.ToString(arg); len(str) > 0 {
			buf.WriteString(str)

			if i < len(elements)-1 {
				buf.WriteString(sep)
			}
		}
	}

	return buf.String()
}

// BeginningOfToday REQUIRE THEM TO DOCUMENT THIS FUNCTION
func BeginningOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// DSN2MAP REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DSN2MAP(dsn string) map[string]string {
	re := regexp.MustCompile("^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?(?:(?P<net>[^\\(]*)(?:\\((?P<addr>[^\\)]*)\\))?)?\\/(?P<dbname>.*?)(?:\\?(?P<params>[^\\?]*))?$")
	match := re.FindStringSubmatch(dsn)

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if len(name) > 0 {
			result[name] = match[i]
		}
	}
	return result
}

// DSN2Publishable REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DSN2Publishable(dsn string) string {
	dsnMap := DSN2MAP(dsn)
	return fmt.Sprintf("%s@%s(%s)/%s?%s",
		dsnMap["user"],
		dsnMap["net"],
		dsnMap["addr"],
		dsnMap["dbname"],
		dsnMap["params"])
}

// Round REQUIRE THEM TO DOCUMENT THIS FUNCTION
func Round(value float64, precision int) float64 {
	exponential := math.Pow10(precision)
	return math.Ceil(value*exponential) / exponential
}

// RandomInt REQUIRE THEM TO DOCUMENT THIS FUNCTION
func RandomInt(bottom, top int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(top-bottom) + bottom
}

// Truncate REQUIRE THEM TO DOCUMENT THIS FUNCTION
func Truncate(s string, i int) (r string) {
	r = s
	if len(s) > i {
		r = s[:i]
	}
	r = strings.TrimSpace(r)
	r = strings.Replace(r, "\n", "", -1)
	r = strings.Replace(r, "    ", "", -1)
	return
}

// Fill merges data from struct instance to another
// By @titpetric suggested in https://scene-si.org/2016/06/01/golang-tips-and-tricks
func Fill(dest interface{}, src interface{}) {
	mSrc := structs.Map(src)
	mDest := structs.Map(dest)
	for key, val := range mSrc {
		if _, ok := mDest[key]; ok {
			structs.New(dest).Field(key).Set(val)
		}
	}
}
