package util

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {

	s = strings.TrimSpace(s)

	if s == "" {
		return s
	}

	result := []string{}

	s = strings.ToLower(s)

	txt := strings.Split(s, " ")

	for _, i := range txt {
		runes := []rune(i)
		runes[0] = unicode.ToUpper(runes[0])
		result = append(result, string(runes))
	}

	return strings.Join(result, " ")

}

func PhoneNumberFormatter(s string) string {

	s = strings.TrimSpace(s)

	if s == "" {
		return s
	}

	runes := []rune(s)

	if string(runes[0:4]) == "+628" {
		return s
	} else if string(runes[0:4]) == "6208" {
		return "+62" + string(runes[3:])
	} else if string(runes[0:3]) == "628" {
		return "+" + s
	} else if string(runes[0:2]) == "08" {
		return "+62" + string(runes[1:])
	}

	return ""

}
