package main

import (
	"fmt"
	"path/filepath"
	"plugin"
	"sync"
	"time"

	"github.com/xellio/modules/pkg/wrapper"
)

func main() {

	for {
		plugins, err := filepath.Glob("./plugins/*.so")
		if err != nil {
			panic(err)
		}

		var wg sync.WaitGroup
		wg.Add(len(plugins))
		for _, filename := range plugins {
			p, err := plugin.Open(filename)
			if err != nil {
				panic(err)
			}

			symbol, err := p.Lookup("New")
			if err != nil {
				panic(err)
			}

			plug, err := symbol.(func() (wrapper.Module, error))()
			if err != nil {
				panic(err)
			}
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				plug.Execute()
			}(&wg)
		}

		wg.Wait()
		fmt.Printf("\n")
		time.Sleep(1 * time.Second)
	}

}
