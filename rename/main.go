package main

import (
	"bufio"
	"fmt"
	"os"

	rename "github.com/CodeAKrome/shed/rename/pkg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		filename := scanner.Text()
		cleanedFilename := rename.CleanFilename(filename)
		fmt.Println(filename + "\t" + cleanedFilename)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
