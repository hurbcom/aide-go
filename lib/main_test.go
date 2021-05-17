package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestRegexpRFC3339(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				input: "2020-10-22T12:46:36Z",
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				input: "2020-11-03T15:24:05-03:00",
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				input: "2020-01-01T16:34:05-03",
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				input: "2020-07-13T19:14:05-0300",
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				input: "2020-11-03T15:24:05Z03:00",
			},
			want: false,
		},
		{
			name: "case6",
			args: args{
				input: "2020-01-01T16:34:05Z03",
			},
			want: false,
		},
		{
			name: "case7",
			args: args{
				input: "2020-07-13T19:14:05Z0300",
			},
			want: false,
		},
		{
			name: "case8",
			args: args{
				input: "2020-11-21T10:50:00.000Z",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := regexpRFC3339.MatchString(tt.args.input); actual != tt.want {
				t.Errorf("TestRegexpRFC3339() %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestGetStringBodyHTTPRequest(t *testing.T) {
	body, _ := json.Marshal(nil)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(body))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "null", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", nil)
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", nil)
	req.Header = nil
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", bytes.NewBuffer([]byte("PLAIN TEXT")))
	req.Header = nil
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "PLAIN TEXT", *actual)
}

func TestGetStringBodyHTTPRequestJSON(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"foo": "bar"})
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(body))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "{\"foo\":\"bar\"}", *actual)
}

func TestGetStringBodyHTTPRequestPlainText(t *testing.T) {
	stringBody := "PLAIN TEXT"
	byteArrayStringBody := []byte(stringBody)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(byteArrayStringBody))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, stringBody, *actual)
}

func TestGetStringBodyHTTPRequestJSONEncoded(t *testing.T) {
	stringBody := `1223ab
{'response':{'code':200}}
0

`
	byteArrayStringBody := []byte(stringBody)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(byteArrayStringBody))
	actual := GetStringBodyHTTPRequestJSON(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "{'response':{'code':200}}", *actual)
}

func TestGetStringBodyHTTPResponseJSON(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	req, _ := http.NewRequest("GET", "http://server.com/bar", nil)
	client := &http.Client{}
	res, _ := client.Do(req)
	actual := GetStringBodyHTTPResponse(res)

	assert.NotNil(t, actual)
	assert.Equal(t, "{\"foo\":\"bar\"}", *actual)
}

func TestGetStringBodyHTTPResponsePlainText(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	stringBody := "PLAIN TEXT"

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		BodyString(stringBody)

	req, err := http.NewRequest("GET", "http://server.com/bar", nil)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	actual := GetStringBodyHTTPResponse(res)

	assert.NotNil(t, actual)
	assert.Equal(t, stringBody, *actual)
}

func TestGetStringBodyHTTPResponseJSONEncoded(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	stringBody := `1223ab
{'response':{'code':200}}
0

`

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		BodyString(stringBody)

	req, err := http.NewRequest("GET", "http://server.com/bar", nil)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	actual := GetStringBodyHTTPResponseJSON(res)

	assert.NotNil(t, actual)
	assert.Equal(t, "{'response':{'code':200}}", *actual)
}

func TestIntegerSliceToStringSlice(t *testing.T) {
	strs := ToStringSlice([]int{1, 2, 3})

	assert.Len(t, strs, 3)
	assert.Equal(t, "1", strs[0])
	assert.Equal(t, "2", strs[1])
	assert.Equal(t, "3", strs[2])
}

func TestInteger64SliceToStringSlice(t *testing.T) {
	strs := ToStringSlice64([]int64{16, 23, 39})

	assert.Len(t, strs, 3)
	assert.Equal(t, "16", strs[0])
	assert.Equal(t, "23", strs[1])
	assert.Equal(t, "39", strs[2])
}

func TestToIntSlice(t *testing.T) {
	actual := ToIntSlice([]string{"6549", "8523", "a"})

	assert.Len(t, actual, 2)
	assert.Equal(t, int(6549), actual[0])
	assert.Equal(t, int(8523), actual[1])
}

func TestToInt64Slice(t *testing.T) {
	actual := ToInt64Slice([]string{"654987", "852369", "a"})

	assert.Len(t, actual, 2)
	assert.Equal(t, int64(654987), actual[0])
	assert.Equal(t, int64(852369), actual[1])
}

func TestStringToStringSlice(t *testing.T) {
	actual := StringToStringSlice("[foo,123,bar,,456,a1b2,,,]")

	assert.Len(t, actual, 5)
	assert.Equal(t, "foo", actual[0])
	assert.Equal(t, "123", actual[1])
	assert.Equal(t, "bar", actual[2])
	assert.Equal(t, "456", actual[3])
	assert.Equal(t, "a1b2", actual[4])

	actual = StringToStringSlice("[[foo,123,[bar],,456,,a1b2]")

	assert.Len(t, actual, 5)
	assert.Equal(t, "foo", actual[0])
	assert.Equal(t, "123", actual[1])
	assert.Equal(t, "bar", actual[2])
	assert.Equal(t, "456", actual[3])
	assert.Equal(t, "a1b2", actual[4])
}

func TestStringToIntSlice(t *testing.T) {
	actual := StringToIntSlice("[foo,123,bar,,456,a1b2,,,]")

	assert.Len(t, actual, 2)
	assert.Equal(t, 123, actual[0])
	assert.Equal(t, 456, actual[1])

	actual = StringToIntSlice("[[foo,123,[bar],,456,,a1b2]")

	assert.Len(t, actual, 2)
	assert.Equal(t, 123, actual[0])
	assert.Equal(t, 456, actual[1])
}

func TestParseInt(t *testing.T) {
	i, err := ParseStringToInt("6549")

	expected := int(6549)

	assert.Empty(t, err)
	assert.IsType(t, expected, i)
	assert.Equal(t, expected, i)
}

func TestParseIntWithEmptyString(t *testing.T) {
	i, err := ParseStringToInt("")

	assert.Equal(t, 0, i)
	assert.Empty(t, err)
}

func TestParseIntInvalidString(t *testing.T) {
	_, err := ParseStringToInt("invalid")

	assert.NotEmpty(t, err)
}

func TestParseInt64(t *testing.T) {
	i, err := ParseStringToInt64("456123789123")

	expected := int64(456123789123)

	assert.Empty(t, err)
	assert.IsType(t, expected, i)
	assert.Equal(t, expected, i)
}

func TestParseInt64WithEmptyString(t *testing.T) {
	i, err := ParseStringToInt64("")

	assert.Equal(t, int64(0), i)
	assert.Empty(t, err)
}

func TestParseInt64InvalidString(t *testing.T) {
	_, err := ParseStringToInt64("invalid")

	assert.NotEmpty(t, err)
}

func TestShouldParseTimeWithYearMonthDayPattern(t *testing.T) {
	date, err := ParseDateYearMonthDay("2000-12-31")
	assert.Nil(t, err)
	assert.False(t, date.IsZero())
	assert.EqualValues(t, 2000, date.Year())
	assert.EqualValues(t, 12, date.Month())
	assert.EqualValues(t, 31, date.Day())
	assert.EqualValues(t, 0, date.Hour())
	assert.EqualValues(t, 0, date.Minute())
	assert.EqualValues(t, 0, date.Second())
}

func TestShouldNotParseTimeWithoutYearMonthDayPattern(t *testing.T) {
	var err error
	_, err = ParseDateYearMonthDay("01-12-2000")
	assert.NotNil(t, err)

	_, err = ParseDateYearMonthDay("01-12-00")
	assert.NotNil(t, err)
}

func TestDiffDays(t *testing.T) {
	duration, err := DiffDays(time.Date(2016, 2, 5, 0, 0, 0, 0, time.UTC), time.Date(2016, 2, 11, 0, 0, 0, 0, time.UTC))
	assert.NotNil(t, duration)
	assert.Nil(t, err)
	assert.Equal(t, int64(6), duration)

	duration, err = DiffDays(time.Date(2016, 2, 20, 0, 0, 0, 0, time.UTC), time.Date(2016, 3, 10, 0, 0, 0, 0, time.UTC))
	assert.NotNil(t, duration)
	assert.Nil(t, err)
	assert.Equal(t, int64(19), duration)

	date1 := time.Time{}
	date2 := time.Time{}
	duration, err = DiffDays(date1, date2)
	assert.Empty(t, duration)
	assert.NotNil(t, err)
}

func TestParseDateStringToTime(t *testing.T) {
	p := func(t time.Time) *time.Time {
		return &t
	}

	type args struct {
		dateString string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{dateString: "2016-01-01"},
			want:    p(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "case2",
			args:    args{dateString: "2016-01-01T00:00:00"},
			want:    p(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "case3",
			args:    args{dateString: "2016-01-01T00:00:00Z"},
			want:    p(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "case4",
			args:    args{dateString: "2016-01-01 00:00:00"},
			want:    p(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "case5",
			args:    args{dateString: "2016-01-01T00:00:00+00:00"},
			want:    p(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			name:    "case6",
			args:    args{dateString: "2016-01-01T00:00:00ABC"},
			want:    nil,
			wantErr: true,
		},
		// zeroed times
		{
			name:    "case7",
			args:    args{dateString: "0000-00-00"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case8",
			args:    args{dateString: "0000-00-00T00:00:00"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case9",
			args:    args{dateString: "0000-00-00T00:00:00Z"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case10",
			args:    args{dateString: "0000-00-00 00:00:00"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case11",
			args:    args{dateString: "0000-00-00T00:00:00+00:00"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "case12",
			args:    args{dateString: "0000-00-00T00:00:00ABC"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateStringToTime(tt.args.dateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateStringToTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want == nil && got != nil {
				t.Errorf("ParseDateStringToTime() = %v, want nil", err)
			}
			if tt.want != nil && got != nil && !(*tt.want).Equal(*got) {
				t.Errorf("ParseDateStringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShouldParseIntToBool(t *testing.T) {
	result1 := ParseIntToBool(0)
	assert.Equal(t, false, result1)

	result2 := ParseIntToBool(1)
	assert.Equal(t, true, result2)

	result3 := ParseIntToBool(2)
	assert.Equal(t, false, result3)

	result4 := ParseIntToBool(345)
	assert.Equal(t, false, result4)
}

func TestShouldParseBoolToString(t *testing.T) {
	result1 := ParseBoolToString(true)
	assert.Equal(t, "1", result1)

	result2 := ParseBoolToString(false)
	assert.Equal(t, "0", result2)
}

func TestShouldCheckStringJSONData(t *testing.T) {
	var s string
	result1 := CheckStringJSONData(s)
	assert.Nil(t, result1)

	result2 := CheckStringJSONData("")
	assert.Nil(t, result2)

	result3 := CheckStringJSONData("test")
	assert.NotNil(t, result3)
	assert.Equal(t, "test", *result3)
}

func TestShouldCheckInt64JSONData(t *testing.T) {
	var i1 int64
	result1 := CheckInt64JSONData(i1)
	assert.Nil(t, result1)

	result2 := CheckInt64JSONData(0)
	assert.Nil(t, result2)

	result3 := CheckInt64JSONData(987654)
	assert.NotNil(t, result3)
	assert.Equal(t, int64(987654), *result3)
}

func TestShouldCheckFloat64JSONData(t *testing.T) {
	var f1 float64
	result1 := CheckFloat64JSONData(f1)
	assert.Nil(t, result1)

	result2 := CheckFloat64JSONData(0)
	assert.Nil(t, result2)

	result3 := CheckFloat64JSONData(0.00)
	assert.Nil(t, result3)

	result4 := CheckFloat64JSONData(9876.54)
	assert.NotNil(t, result4)
	assert.Equal(t, float64(9876.54), *result4)
}

func TestShouldReturnOnlyNumbers(t *testing.T) {
	s1 := "61.225.412/0001-14aA"

	result := *GetOnlyNumbers(&s1)
	assert.Equal(t, "61225412000114", result)
}

func TestShouldReturnOnlyNumbersOrSpecial(t *testing.T) {
	s1 := "+55 (21) 98765-4321"

	result := *GetOnlyNumbersOrSpecial(&s1, "+")
	assert.Equal(t, "+5521987654321", result)
}

func TestShouldReturnOnlyNumbersOrSpecial1(t *testing.T) {
	s1 := "+55 (21) 98765-4321"

	result := *GetOnlyNumbersOrSpecial(&s1, "+()")
	assert.Equal(t, "+55(21)987654321", result)
}

func TestShouldReturnNilForNilInput(t *testing.T) {
	var s1 string

	result := *GetOnlyNumbers(&s1)
	assert.Equal(t, s1, result)
}

func TestShouldReturnNilForNilInput1(t *testing.T) {
	var s1 string

	result := *GetOnlyNumbersOrSpecial(&s1, "+")
	assert.Equal(t, s1, result)
}

func TestParseIntOrReturnZero(t *testing.T) {
	stg := "1"
	expected := 1

	assert.Equal(t, expected, ParseIntOrReturnZero(stg))
}

func TestParseIntOrReturnZeroFail(t *testing.T) {
	stg := "a"
	expected := 0

	assert.Equal(t, expected, ParseIntOrReturnZero(stg))
}

func TestParseIntOrReturnZeroWithNumberOnString(t *testing.T) {
	stg := "a123"
	expected := 0

	assert.Equal(t, expected, ParseIntOrReturnZero(stg))
}

func TestIsArray(t *testing.T) {
	actual := IsArray([]string{"foo", "bar"})
	assert.Equal(t, true, actual)

	actual = IsArray([]int{65485, 19734})
	assert.Equal(t, true, actual)

	actual = IsArray([]int64{65485, 19734})
	assert.Equal(t, true, actual)

	actual = IsArray(nil)
	assert.Equal(t, false, actual)

	actual = IsArray(65485)
	assert.Equal(t, false, actual)

	actual = IsArray("foo")
	assert.Equal(t, false, actual)

	actual = IsArray(false)
	assert.Equal(t, false, actual)
}

func TestIsString(t *testing.T) {
	actual := IsString("foo")
	assert.Equal(t, true, actual)

	actual = IsString("")
	assert.Equal(t, true, actual)

	actual = IsString("123")
	assert.Equal(t, true, actual)

	actual = IsString("123.456")
	assert.Equal(t, true, actual)

	actual = IsString("true")
	assert.Equal(t, true, actual)

	actual = IsString([]int{1, 2})
	assert.Equal(t, false, actual)

	actual = IsString([]string{"a", "b"})
	assert.Equal(t, false, actual)

	actual = IsString(nil)
	assert.Equal(t, false, actual)

	actual = IsString(123)
	assert.Equal(t, false, actual)

	actual = IsString(123.456)
	assert.Equal(t, false, actual)

	actual = IsString(false)
	assert.Equal(t, false, actual)
}

func TestIsPointer(t *testing.T) {
	var pStr *string
	actual := IsPointer(pStr)
	assert.Equal(t, true, actual)

	var pInt *int
	actual = IsPointer(pInt)
	assert.Equal(t, true, actual)

	var pInt64 *int64
	actual = IsPointer(pInt64)
	assert.Equal(t, true, actual)

	var pFloat64 *float64
	actual = IsPointer(pFloat64)
	assert.Equal(t, true, actual)

	var pInter *interface{}
	actual = IsPointer(pInter)
	assert.Equal(t, true, actual)

	var pSlice *[]string
	actual = IsPointer(pSlice)
	assert.Equal(t, true, actual)

	var strVar string
	actual = IsPointer(strVar)
	assert.Equal(t, false, actual)

	var intVar int
	actual = IsPointer(intVar)
	assert.Equal(t, false, actual)

	var int64Var int64
	actual = IsPointer(int64Var)
	assert.Equal(t, false, actual)

	var float64Var float64
	actual = IsPointer(float64Var)
	assert.Equal(t, false, actual)

	var interVar interface{}
	actual = IsPointer(interVar)
	assert.Equal(t, false, actual)

	var sliceValue []string
	actual = IsPointer(sliceValue)
	assert.Equal(t, false, actual)
}

func TestJoin(t *testing.T) {
	actual := Join(", ", 654321987, "bar", 654.654)
	assert.Equal(t, `654321987, bar, 654.654`, actual)

	actual = Join(", ", int64(654321987), "bar")
	assert.Equal(t, `654321987, bar`, actual)

	actual = Join(", ", int64(654321987), int64(52354))
	assert.Equal(t, `654321987, 52354`, actual)

	actual = Join(", ", "foo")
	assert.Equal(t, `foo`, actual)

	actual = Join(", ", []string{"foo", "bar"})
	assert.Equal(t, `foo, bar`, actual)

	actual = Join(", ", []int{65485, 19734})
	assert.Equal(t, `65485, 19734`, actual)

	actual = Join("_", []string{"foo", "bar", ""}, nil, 123)
	assert.Equal(t, `foo_bar_123`, actual)

	var pStr *string
	str := "foo"

	actual = Join(", ", 654321987, nil, 654.654, "", pStr, &str)
	assert.Equal(t, `654321987, 654.654, foo`, actual)
}

func TestBeginningOfToday(t *testing.T) {
	today := BeginningOfToday()
	assert.Equal(t, today.Year(), time.Now().Year())
	assert.Equal(t, today.Month(), time.Now().Month())
	assert.Equal(t, today.Day(), time.Now().Day())
	assert.Equal(t, today.Hour(), 0)
	assert.Equal(t, today.Minute(), 0)
	assert.Equal(t, today.Second(), 0)
}

func TestBeginningOfTodayIn(t *testing.T) {
	loc, _ := time.LoadLocation("Pacific/Fakaofo")
	today := BeginningOfTodayIn(loc)
	assert.Equal(t, today.Year(), time.Now().Year())
	assert.Equal(t, today.Month(), time.Now().Month())
	assert.Equal(t, today.Day(), time.Now().Day()+1)
	assert.Equal(t, today.Hour(), 0)
	assert.Equal(t, today.Minute(), 0)
	assert.Equal(t, today.Second(), 0)
}

func TestShouldRemoveNanoseconds(t *testing.T) {
	expected := time.Date(2016, time.September, 20, 18, 49, 15, 0, time.UTC)

	date := time.Date(2016, time.September, 20, 18, 49, 15, 999999999, time.UTC)

	actual, err := RemoveNanoseconds(date)
	assert.Nil(t, err)
	assert.EqualValues(t, expected.Year(), actual.Year())
	assert.EqualValues(t, expected.Month(), actual.Month())
	assert.EqualValues(t, expected.Day(), actual.Day())
	assert.EqualValues(t, expected.Hour(), actual.Hour())
	assert.EqualValues(t, expected.Minute(), actual.Minute())
	assert.EqualValues(t, expected.Second(), actual.Second())
	assert.EqualValues(t, expected.Local(), actual.Local())
	assert.EqualValues(t, expected.Nanosecond(), actual.Nanosecond(), "need to be zeroed")
}

func TestDSN2MAP(t *testing.T) {
	dsn := "MyUser:MyPassword@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"
	actual := DSN2MAP(dsn)

	assert.Equal(t, "MyUser", actual["user"])
	assert.Equal(t, "MyPassword", actual["passwd"])
	assert.Equal(t, "tcp", actual["net"])
	assert.Equal(t, "MyDatabaseHost.net:3306", actual["addr"])
	assert.Equal(t, "MyDatabaseName", actual["dbname"])
	assert.Equal(t, "param1=1&param2=2", actual["params"])
}

func TestDSN2Publishable(t *testing.T) {
	dsn := "MyUser:MyPassword@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"
	expected := "MyUser@tcp(MyDatabaseHost.net:3306)/MyDatabaseName?param1=1&param2=2"

	assert.Equal(t, expected, DSN2Publishable(dsn))
}

func TestGetByteArrayAndBufferFromRequestBody(t *testing.T) { t.Skip("Implement this test") }
func TestGetOnlyNumbers(t *testing.T)                       { t.Skip("Implement this test") }
func TestGetOnlyNumbersOrSpecial(t *testing.T)              { t.Skip("Implement this test") }
func TestParseStringToBool(t *testing.T)                    { t.Skip("Implement this test") }
func TestParseStringToInt(t *testing.T)                     { t.Skip("Implement this test") }
func TestParseStringToInt64(t *testing.T)                   { t.Skip("Implement this test") }
func TestToStringSlice64(t *testing.T)                      { t.Skip("Implement this test") }

func TestRound(t *testing.T) {
	assert.Equal(t, 1.2, Round(float64(1.2), 2))
	assert.Equal(t, 1.23, Round(float64(1.23), 2))
	assert.Equal(t, 1.24, Round(float64(1.233), 2))
	assert.Equal(t, 1.24, Round(float64(1.237), 2))
	assert.Equal(t, 1234.56, Round(float64(1234.56), 2))
	assert.Equal(t, 1234.567, Round(float64(1234.567), 3))
	assert.Equal(t, 1234.568, Round(float64(1234.5674), 3))
	assert.Equal(t, 1234.568, Round(float64(1234.5678), 3))
}

func TestRandomInt(t *testing.T) {
	a := RandomInt(1, 9999)
	b := RandomInt(1, 9999)
	c := RandomInt(1, 9999)
	d := RandomInt(1, 9999)
	e := RandomInt(1, 9999)

	assert.NotEqual(t, a, b)
	assert.NotEqual(t, a, c)
	assert.NotEqual(t, a, d)
	assert.NotEqual(t, a, e)
	assert.NotEqual(t, b, c)
	assert.NotEqual(t, b, d)
	assert.NotEqual(t, b, e)
	assert.NotEqual(t, c, d)
	assert.NotEqual(t, c, e)
	assert.NotEqual(t, d, e)
}

func TestTruncate(t *testing.T) {
	x := "123456789abcdef"
	assert.Equal(t, x, Truncate(x, 16))
	assert.Equal(t, x[:15], Truncate(x, 15))
	assert.Equal(t, "12", Truncate(x, 2))

	expected := "{123456789abcdef}"
	json1 := `{
        123456789abcdef
    }`
	assert.Equal(t, expected, Truncate(json1, len(json1)))
}

func TestFill(t *testing.T) {
	a := struct {
		ID      int
		Name    string
		IsAdmin bool
	}{}

	b := struct {
		Name string
	}{}

	b.Name = "Bobby"

	Fill(&a, b)

	assert.Equal(t, "Bobby", a.Name)
}

func TestParseStringToFloat64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Should parse ok",
			args: args{
				s: "123.123",
			},
			want: 123.123,
		},
		{
			name: "Should parse with zero value",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "should parse with zero",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "Should return an error",
			args: args{
				s: "abc",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringToFloat64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringTFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseStringTFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
