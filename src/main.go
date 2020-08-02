package main

import (
	"flag"
	"fmt"
	"os"
	"pdfreader/lib/pdfreader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: pdfreader --file-path=input.pdf\n")
		os.Exit(1)
	}

	var filePath string
	flag.StringVar(&filePath, "file-path", "", "")
	flag.Parse()

	reader := pdfreader.Pdf{Path: filePath}
	content, err := reader.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(content)
}
