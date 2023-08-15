package formatter

import "strings"

func FormatFrom(from string) string {
	return strings.ReplaceAll(from, "\\", "/")
}
