package app

import (
	"fmt"
	"os"

	"github.com/zu1k/proxypool/config"
	"github.com/zu1k/proxypool/getter"
)

var Getters = make([]getter.Getter, 0)

func InitConfigAndGetters(path string) {
	c, err := config.Parse(path)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	if c == nil {
		fmt.Println("Error: no sources")
		os.Exit(2)
	}
	InitGetters(c.Sources)
}

func InitGetters(sources []config.Source) {
	Getters = make([]getter.Getter, 0)
	for _, source := range sources {
		g := getter.NewGetter(source.Type, source.Options)
		if g != nil {
			Getters = append(Getters, g)
			fmt.Println("init getter:", source.Type, source.Options)
		}
	}
	fmt.Println("Getter count:", len(Getters))
}
