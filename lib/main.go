package lib

import (
	"bytes"
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
	"github.com/hurbcom/aide-go/constraints"
	"github.com/spf13/cast"
)

// ToStringSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func IntSliceToStringSlice[T constraints.SignedInteger](intSlice []T) (stringSlice []string) {
	for _, v := range intSlice {
		stringSlice = append(stringSlice, strconv.FormatInt(int64(v), 10))
	}
	return stringSlice
}

// ToIntSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func StringSliceToIntSlice[T constraints.SignedInteger](stringSlice []string) (intSlice []T) {
	for _, i := range stringSlice {
		v, err := ParseStringToInt[T](i)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, T(v))
	}
	return intSlice
}

// StringToStringSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func StringToStringSlice(s string) []string {
	stringSlice := []string{}
	if len(s) == 0 {
		return []string{}
	}

	s1 := regexpCommaAlphaNum.ReplaceAllString(s, "")
	if len(s1) == 0 {
		return []string{}
	}

	s2 := strings.Split(s1, ",")
	if len(s2) == 0 {
		return []string{}
	}

	for _, s3 := range s2 {
		if len(s3) > 0 {
			stringSlice = append(stringSlice, s3)
		}
	}

	return stringSlice
}

// StringToIntSlice REQUIRE THEM TO DOCUMENT THIS FUNCTION
func StringToIntSlice[T constraints.SignedInteger](s string) []T {
	if len(s) == 0 {
		return []T{}
	}

	sl := StringToStringSlice(s)
	if len(sl) == 0 {
		return []T{}
	}

	return StringSliceToIntSlice[T](sl)
}

// ParseStringToInt64 REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseStringToInt[T constraints.SignedInteger](s string) (T, error) {
	if s == "" {
		return 0, nil
	}

	v, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return T(v), nil 
}

// ParseIntToBool REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ParseIntToBool[T constraints.SignedInteger](i T) bool {
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

	headers, err := httputil.DumpRequest(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpRequest(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPRequestJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPRequestJSON(r *http.Request) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpRequest(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpRequest(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

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

	headers, err := httputil.DumpResponse(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpResponse(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPResponseJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPResponseJSON(r *http.Response) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpResponse(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpResponse(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

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
func ParseIntOrReturnZero[T constraints.SignedInteger](s string) T {
	integer, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return T(integer)
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

// DSN2MAP REQUIRE THEM TO DOCUMENT THIS FUNCTION
func DSN2MAP(dsn string) map[string]string {
	re := regexp.MustCompile("^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?(?:(?P<net>[^\\(]*)(?:\\((?P<addr>[^\\)]*)\\))?)?\\/(?P<dbname>.*?)(?:\\?(?P<params>[^\\?]*))?$")
	match := re.FindStringSubmatch(dsn)

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if len(name) > 0 && i < len(match) {
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

// ParseStringToFloat64 parse the string to float64
func ParseStringToFloat64(s string) (float64, error) {
	if s == "" || s == "0" {
		return float64(0), nil
	}
	return strconv.ParseFloat(s, 64)
}
