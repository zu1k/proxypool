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
