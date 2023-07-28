package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	from := strings.ReplaceAll(os.Args[1], "\\", "/")
	fromParts := strings.Split(from, "/")
	wxid := fromParts[len(fromParts)-3]

	needDecrypt := strings.HasPrefix(wxid, "wx")
	if needDecrypt {
		fmt.Println("wxid", wxid)
		err := Decrypt(from, wxid)
		if err != nil {
			fmt.Println(err)
			return
		}
		from += DecryptTo
	}

	err := Unpack(from)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
