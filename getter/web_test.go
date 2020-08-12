package getter

import (
	"fmt"
	"testing"
)

func TestWebLucnOrg_Get(t *testing.T) {
	fmt.Println(WebLucnOrg{}.Get())
}

func TestWebFreessrXyz_Get(t *testing.T) {
	fmt.Println(WebFreessrXyz{}.Get())
}

func TestWebFuzz_Get(t *testing.T) {
	fmt.Println(NewWebFuzz("https://merlinblog.xyz/wiki/freess.html").Get())
}

func TestSubscribe_Get(t *testing.T) {
	fmt.Println(NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/v2ray").Get())
	fmt.Println(NewSubscribe("https://raw.githubusercontent.com/ssrsub/ssr/master/ssrsub").Get())
}
