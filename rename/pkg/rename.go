package rename

import (
	"path/filepath"
	"regexp"
	"strings"
)

func parseFilePath(filePath string) (dir string, base string, ext string) {
	// Get directory and base filename
	dir, base = filepath.Split(filePath)

	// Remove trailing slash from directory if present
	dir = strings.TrimSuffix(dir, "/")

	// Get file extension (if any)
	ext = filepath.Ext(base)

	// Get filename without extension
	base = strings.TrimSuffix(base, ext)

	return dir, base, ext
}

// Example filename:
// Downloads/Trading Places (1983) [2160p] [4K] [BluRay] [5.1] [YTS.MX]/Trading.Places.1983.2160p.4K.BluRay.x265.10bit.AAC5.1-[YTS.MX].mkv
func CleanFilename(filePath string) (file string) {
	// Remove leading and trailing spaces
	// filename = strings.TrimSpace(file)
	// Split the filename into parts
	_, base, ext := parseFilePath(filePath)

	// Replace invalid characters with underscores, except for years
	re := regexp.MustCompile(`.*(19|20)\d{2}`) // Match 4-digit years
	match := re.FindString(base)

	// Surround date with parens
	re = regexp.MustCompile(`(19|20)\d{2}`)
	match = re.ReplaceAllString(match, "($0)")

	// Delete invalid characters
	re = regexp.MustCompile(`[^A-Za-z0-9\-\(\)\. ]`)
	filename := re.ReplaceAllString(match, "")

	// Replace periods with spaces
	re = regexp.MustCompile(`\.`)
	filename = re.ReplaceAllString(filename, " ")
	file = filename + ext

	return file
}
