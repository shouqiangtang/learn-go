package main

import (
	"fmt"
	"learn-go/maze"
	"os"
)

func main() {
	data, err := maze.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)

	start := maze.Point{0, 0}
	end := maze.Point{5, 4}
	footPrints, err := maze.Walk(data, start, end)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(footPrints)
}
