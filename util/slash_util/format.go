package slash_util

import "strings"

func Format(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
