package main

import (
	"fmt"
	"github.com/zhuweiyou/wxapkg/decrypter"
	"github.com/zhuweiyou/wxapkg/unpacker"
	"os"
)

func main() {
	defer func() {
		os.Stdin.Read(make([]byte, 1))
	}()

	if len(os.Args) < 2 {
		fmt.Println("请查看 readme.md 使用说明")
		return
	}

	from := decrypter.FormatFrom(os.Args[1])
	fmt.Println("from", from)

	wxid, needDecrypt := decrypter.GetWxid(from)
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
