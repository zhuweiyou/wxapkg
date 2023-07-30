package decrypter

import "strings"

func GetWxid(from string) (string, bool) {
	parts := strings.Split(from, "/")

	var wxid string
	wxidIndex := len(parts) - 3
	if wxidIndex >= 0 {
		wxid = parts[wxidIndex]
	}

	return wxid, strings.HasPrefix(wxid, "wx")
}
