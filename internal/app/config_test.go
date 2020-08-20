package app

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/zu1k/proxypool/config"
	"github.com/zu1k/proxypool/pkg/getter"
)

func TestConfigFile(t *testing.T) {
	c, err := config.Parse("../source.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	if c == nil {
		t.Error(errors.New("no sources"))
		return
	}
	for idx, source := range c.Sources {
		g, err := getter.NewGetter(source.Type, source.Options)
		if err != nil {
			t.Error(err, idx)
			fmt.Println(source)
			return
		}
		if g == nil {
			t.Error(errors.New("getter is nil:" + strconv.Itoa(idx)))
			fmt.Println(source)
			return
		}
	}
}
