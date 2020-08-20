package tool

import (
	"regexp"
	"unicode"
)

var hanRe = regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]")

func ContainChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (hanRe.MatchString(string(r))) {
			return true
		}
	}
	return false
}
