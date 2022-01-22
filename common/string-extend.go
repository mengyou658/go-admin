package common

import "strings"

func IsEmpty(str string) bool {
	return str == "" || strings.TrimSpace(str) == ""
}
