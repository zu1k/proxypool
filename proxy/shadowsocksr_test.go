package proxy

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGrepSSRLinkFromString(t *testing.T) {
	fmt.Println(GrepSSRLinkFromString("ssr://abcssr://def ssr://126ssr://789 ssr://abv123"))
}

var link = "ssr://MTcyLjEwNC4xMjcuMjA4OjgwOTc6b3JpZ2luOmFlcy0yNTYtY2ZiOnBsYWluOlpVbFhNRVJ1YXpZNU5EVTBaVFp1VTNkMWMzQjJPVVJ0VXpJd01YUlJNRVEvP3JlbWFya3M9NTctNzVhS1o1WVdhWm1GdWNXbGhibWRrWVc1bkxtTnZiUSY9UUZOVFVsTlZRaTNtbDZYbW5LeEJNVEV0NUx1WTZMUzVVMU5TNW82bzZJMlFPblF1WTI0dlJVZEtTWGx5YkEmcHJvdG9wYXJhbT1kQzV0WlM5VFUxSlRWVUkmb2Jmc3BhcmFtPTVMdVk2TFM1VTFOUzVyT281WWFNT21oMGRIQTZMeTkwTG1OdUwwVkhTa2w1Y213"

func TestParseSSRLink(t *testing.T) {
	fmt.Println(ParseSSRLink(link))
}

func TestBase64(t *testing.T) {
	remarks := "57-75aKZ5YWaZmFucWlhbmdkYW5nLmNvbQ"
	var dstbytes []byte
	dstbytes, err := base64.RawURLEncoding.DecodeString(remarks)
	fmt.Println(string(dstbytes), err)
	fmt.Println(dstbytes)

	dstbytes, err = base64.RawStdEncoding.DecodeString(remarks)
	fmt.Println(string(dstbytes), err)

	dstbytes, err = base64.URLEncoding.DecodeString(remarks)
	fmt.Println(string(dstbytes), err)

	dstbytes, err = base64.StdEncoding.DecodeString(remarks)
	fmt.Println(string(dstbytes), err)

	a := "中文"
	fmt.Println(a, string([]byte(a)))
	b := base64.URLEncoding.EncodeToString([]byte(a))
	fmt.Println(b)
	c, err := base64.URLEncoding.DecodeString(b)
	fmt.Println(string(c), err)
}
