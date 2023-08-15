package formatter

import "strings"

func FormatSlash(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
