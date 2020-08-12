package app

import (
	"fmt"

	"github.com/zu1k/proxypool/config"
	"github.com/zu1k/proxypool/getter"
)

var Getters = make([]getter.Getter, 0)

func InitGetters(sources []config.Source) {
	for _, source := range sources {
		g := getter.NewGetter(source.Type, source.Options)
		if g != nil {
			Getters = append(Getters, g)
			fmt.Println("init getter:", source.Type, source.Options)
		}
	}
	fmt.Println("Getter count:", len(Getters))
}
