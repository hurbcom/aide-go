package v4

import (
	"fmt"
	"regexp"
)

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

func DSN2Publishable(dsn string) string {
	dsnMap := DSN2MAP(dsn)
	return fmt.Sprintf("%s@%s(%s)/%s?%s",
		dsnMap["user"],
		dsnMap["net"],
		dsnMap["addr"],
		dsnMap["dbname"],
		dsnMap["params"])
}
