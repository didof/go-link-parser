package main

import (
	"fmt"
	"html-link-parser/cmd"
	"html-link-parser/link"
	"os"
)

func main() {
	filename := cmd.ParseArgs()

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", links)
}
