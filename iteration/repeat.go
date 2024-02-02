package iteration

import "strings"

const (
	repeatCount = 5
)

func Repeat(character string, times int) string {
	// var repeated string

	// for i := 0; i < times; i++ {
	// 	repeated += character
	// }

	// return repeated
	return strings.Repeat(character, times)
}
