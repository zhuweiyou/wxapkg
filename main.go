package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		os.Stdin.Read(make([]byte, 1))
	}()

	if len(os.Args) < 2 {
		fmt.Println("请查看 readme.md 使用说明")
		return
	}

	from := os.Args[1]
	fmt.Println("from", from)

	from = strings.ReplaceAll(from, "\\", "/")
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
