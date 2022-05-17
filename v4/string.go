package aidego

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

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

type Stringer interface {
	String() string
}

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
