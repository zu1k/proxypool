package provider

import (
	"fmt"
	"testing"

	"github.com/zu1k/proxypool/getter"
)

func TestClash_Provide(t *testing.T) {
	a := getter.NewTGSsrlistGetter(200).Get()

	clash := Clash{Proxies: a}
	fmt.Println(clash.Provide())

	//data, _ := yaml.Marshal(a)
	//fmt.Println(string(data))
}
