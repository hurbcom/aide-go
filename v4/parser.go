package aidego

import (
	"strconv"
	"strings"
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

func StringToIntSlice(s string) []int {
	if len(s) == 0 {
		return []int{}
	}

	sl := StringToStringSlice(s)
	if len(sl) == 0 {
		return []int{}
	}

	return ToIntSlice(sl)
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

func ParseIntToBool(i int) bool {
	return i == 1
}

func ParseStringToBool(s string) bool {
	return s == "1"
}

func ParseBoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func ParseIntOrReturnZero(s string) int {
	integer, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(integer)
}

// ParseStringToFloat64 parse the string to float64
func ParseStringToFloat64(s string) (float64, error) {
	if s == "" || s == "0" {
		return float64(0), nil
	}
	return strconv.ParseFloat(s, 64)
}
