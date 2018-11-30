package dataparser

import "strings"

func ParseHeaderline(hl string) []string {
	return strings.Split(hl, ",")
}
