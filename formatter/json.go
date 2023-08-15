package formatter

import (
	"encoding/json"
	"strings"
)

func FormatJsonString(s string) []byte {
	bytes := []byte(s)
	var data any
	json.Unmarshal(bytes, &data)
	return FormatJson(data)
}

func FormatJson(data any) []byte {
	formatted, _ := json.MarshalIndent(data, "", strings.Repeat(" ", 4))
	return formatted
}
