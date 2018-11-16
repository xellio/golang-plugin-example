package main

import (
	"fmt"

	"github.com/xellio/modules/pkg/wrapper"
)

const (
	// ModuleName ...
	ModuleName = "FIRST"
)

//
// ModuleStruct is an instance of wrapper.Module
//
type ModuleStruct struct{}

//
// New creates a new ModuleStruct
//
func New() (wrapper.Module, error) {
	return ModuleStruct{}, nil
}

//
// Execute holds the logic for the plugin
//
func (s ModuleStruct) Execute() {
	fmt.Printf("Hello from %s...\n", ModuleName)
}
