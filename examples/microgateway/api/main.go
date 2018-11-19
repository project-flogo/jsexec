package main

import (
	"github.com/project-flogo/core/engine"
	tests "github.com/project-flogo/jsexec/examples/microgateway"
)

func main() {
	e, err := tests.Example()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
