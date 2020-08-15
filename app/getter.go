package app

import (
	"errors"
	"fmt"

	"github.com/zu1k/proxypool/config"
	"github.com/zu1k/proxypool/getter"
)

var Getters = make([]getter.Getter, 0)

func InitConfigAndGetters(path string) (err error) {
	c, err := config.Parse(path)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("no sources")
	}
	InitGetters(c.Sources)
	return nil
}

func InitGetters(sources []config.Source) {
	Getters = make([]getter.Getter, 0)
	for _, source := range sources {
		g, err := getter.NewGetter(source.Type, source.Options)
		if err == nil && g != nil {
			Getters = append(Getters, g)
			fmt.Println("init getter:", source.Type, source.Options)
		}
	}
	fmt.Println("Getter count:", len(Getters))
}
