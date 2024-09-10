package main

import (
	"fmt"

	rename "github.com/CodeAKrome/shed/rename/pkg"
)

func main() {
	filename := "my file.with spaces and ?chars 1957.mkv"
	cleanedFilename := rename.CleanFilename(filename)
	fmt.Println(cleanedFilename)
}
