package main

import (
	"flag"
	"fmt"
	"os"

	"dirsize/internal/scanner"
)

func main() {
	path := flag.String("path", "", "The path to scan")
	human := flag.Bool("human", false, "Human-readable output")
	fullTree := flag.Bool("tree", false, "Show full directory tree")

	// Parse the flags from the command line
	flag.Parse()

	if path == nil || *path == "" {
		fmt.Println("Error: path is required")
		os.Exit(1)
	}

	fmt.Println("Parsing path:", path, "human:", human, "...")

	scanner := scanner.NewScanner(*path)
	tree, err := scanner.Scan()

	if err != nil {
		fmt.Println("Error scanning directory:", err)
		os.Exit(1)
	}

	tree.PrintTree(*human, *fullTree)
}
