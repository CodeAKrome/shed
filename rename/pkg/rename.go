package rename

import (
	"path/filepath"
	"regexp"
	"strings"
)

func CleanFilename(filename string) string {
	// Remove leading and trailing spaces
	filename = strings.TrimSpace(filename)

	// Replace invalid characters with underscores, except for years
	re := regexp.MustCompile(`([12]\d{3})`) // Match 4-digit years
	filename = re.ReplaceAllStringFunc(filename, func(match string) string {
		return "(" + match + ")"
	})

	extension := filepath.Ext(filename)
	if extension != "" {
		filename = filename[:len(filename)-len(extension)]
	}

	// Delete invalid characters
	re2 := regexp.MustCompile(`[^A-Za-z0-9\-\(\)\. ]`)
	filename = re2.ReplaceAllString(filename, "")

	// Replace periods with spaces
	re = regexp.MustCompile(`\.`) // Match periods not preceded by a parenthesis
	filename = re.ReplaceAllString(filename, " ")

	return filename + extension
}
