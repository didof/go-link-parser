package cmd

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() string {
	filename := flag.String("filename", "mock.html", "Path of the HTML file to be parsed")
	flag.Parse()

	if *filename == "" {
		fmt.Fprint(os.Stderr, "Error: empty path\n")
	}

	return *filename
}
