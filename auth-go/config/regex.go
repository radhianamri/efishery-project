package config

import "strings"

const (
	//RegexPhoneNumber example: +6209382929
	RegexPhoneNumber = "^[+][0-9]{7,15}$"
)

var (
	phoneReplacer = strings.NewReplacer(" ", "", "-", "")
)

//SimplifyPhoneNumber creates a simplified compact string from given input
func SimplifyPhoneNumber(s string) string {
	return phoneReplacer.Replace(s)
}
