package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		{
			name:    "case13",
			args:    args{dateString: "2021-08-19 01:27:40.569972"},
			want:    p(time.Date(2021, 8, 19, 1, 27, 40, 569972000, time.UTC)),
			wantErr: false,
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
