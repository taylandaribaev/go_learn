package main

import (
	"fmt"
	"os"

	"dirsize/internal/formatter"
	"dirsize/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dirsize <path> [-h]")
		os.Exit(1)
	}

	path := os.Args[1]
	human := len(os.Args) > 2 && os.Args[2] == "-h"

	fmt.Println("path:", path, "human:", human)

	scanner := scanner.NewScanner(path)
	size, err := scanner.Scan()
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		os.Exit(1)
	}
	fmt.Printf("Total size: %s\n", formatter.FormatSizeExternal(size, human))
}
