package main

import (
	"fmt"
	"learning-go/HashHive/hasher"
	"os"
)

func main() {
	start := os.Args[1]
	outputFile := "file_hashes.csv"

	err := hasher.WalkAndHash(start, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hashes saved to %s\n", outputFile)
	}
}
