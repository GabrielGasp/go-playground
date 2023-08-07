package replace

import (
	"regexp"
	"strings"
)

func ReplaceWithRegex(str string, params map[string]string) string {
	pattern := regexp.MustCompile("{{_p_(.+?)}}")

	replacedSTR := pattern.ReplaceAllStringFunc(str, func(placeholder string) string {
		paramKey := pattern.FindStringSubmatch(placeholder)[1]

		if value, ok := params[paramKey]; ok {
			return value
		}

		return placeholder
	})

	return replacedSTR
}

func ReplaceWithFor(str string, params map[string]string) string {
	replacedSTR := str

	for key, value := range params {
		replacedSTR = strings.ReplaceAll(replacedSTR, "{{_p_"+key+"}}", value)
	}

	return replacedSTR
}
