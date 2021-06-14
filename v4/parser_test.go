package v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
