package json_util

import (
	"encoding/json"
	"strings"
)

func FormatFromString(s string) []byte {
	bytes := []byte(s)
	var data any
	json.Unmarshal(bytes, &data)
	return Format(data)
}

func Format(data any) []byte {
	formatted, _ := json.MarshalIndent(data, "", strings.Repeat(" ", 4))
	return formatted
}
