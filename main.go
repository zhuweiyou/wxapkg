package main

import (
	"fmt"
	"github.com/zhuweiyou/wxapkg/decrypter"
	"github.com/zhuweiyou/wxapkg/unpacker"
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
	fromPaths := strings.Split(from, "/")

	var wxid string
	wxidIndex := len(fromPaths) - 3
	if wxidIndex >= 0 {
		wxid = fromPaths[wxidIndex]
	}
	needDecrypt := strings.HasPrefix(wxid, "wx")
	if needDecrypt {
		fmt.Println("wxid", wxid)
		err := decrypter.DefaultDecrypt(from, wxid)
		if err != nil {
			fmt.Println(err)
			return
		}
		from += decrypter.DefaultDecryptTo
	}

	err := unpacker.Unpack(from)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
