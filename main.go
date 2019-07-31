package main

import (
	"fmt"

	"github.com/rantav/go-archetype/transformer"
	"github.com/rantav/go-archetype/types"
)

func main() {
	fmt.Println("Hello")
	transformations, err := transformer.Read("transformations.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("transformations: ", transformations)
	source := types.Path(".")
	destination := types.Path("/tmp/transformed")
	err = transformer.Transform(source, destination, *transformations)
	if err != nil {
		panic(err)
	}
}
