package app

import (
	"errors"
	"fmt"

	"github.com/zu1k/proxypool/internal/cache"

	"github.com/ghodss/yaml"

	"github.com/zu1k/proxypool/config"
	"github.com/zu1k/proxypool/pkg/getter"
)

var Getters = make([]getter.Getter, 0)

func InitConfigAndGetters(path string) (err error) {
	err = config.Parse(path)
	if err != nil {
		return
	}
	if s := config.Config.SourceFiles; len(s) == 0 {
		return errors.New("no sources")
	} else {
		initGetters(s)
	}
	return
}

func initGetters(sourceFiles []string) {
	Getters = make([]getter.Getter, 0)
	for _, path := range sourceFiles {
		data, err := config.ReadFile(path)
		if err != nil {
			fmt.Errorf("Init SourceFile Error: %s\n", err.Error())
			continue
		}
		sourceList := make([]config.Source, 0)
		err = yaml.Unmarshal(data, &sourceList)
		if err != nil {
			fmt.Errorf("Init SourceFile Error: %s\n", err.Error())
			continue
		}
		for _, source := range sourceList {
			g, err := getter.NewGetter(source.Type, source.Options)
			if err == nil && g != nil {
				Getters = append(Getters, g)
				fmt.Println("init getter:", source.Type, source.Options)
			}
		}
	}
	fmt.Println("Getter count:", len(Getters))
	cache.GettersCount = len(Getters)
}
