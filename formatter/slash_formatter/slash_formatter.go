package slash_formatter

import "strings"

func Format(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
