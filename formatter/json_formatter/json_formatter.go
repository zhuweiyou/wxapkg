package json_formatter

import (
	"encoding/json"
	"strings"
)

func FromString(s string) []byte {
	bytes := []byte(s)
	var data any
	json.Unmarshal(bytes, &data)
	return From(data)
}

func From(data any) []byte {
	formatted, _ := json.MarshalIndent(data, "", strings.Repeat(" ", 4))
	return formatted
}
