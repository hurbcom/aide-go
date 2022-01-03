package aidego

import (
	"fmt"
	"regexp"
	"strings"
)

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
