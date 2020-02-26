package main

import (
	"configure"
	"fmt"
)

var (
	conf = configure.New()
	name = conf.String("name", "Harrison", "The name you want to greet")
)

func init() {
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewFlag())
}

func main() {
	conf.Parse()
	fmt.Printf("Hello, %v\n", *name)
}
