package splitter

import (
	"encoding/json"
	"strings"
)

func FormatJson(jsonBytes []byte) []byte {
	var data interface{}
	json.Unmarshal(jsonBytes, &data)
	formatted, _ := json.MarshalIndent(data, "", strings.Repeat(" ", 4))
	return formatted
}
