package main

import (
	"fmt"
	"github.com/zhuweiyou/wxapkg/decrypter"
	"github.com/zhuweiyou/wxapkg/unpacker"
	"github.com/zhuweiyou/wxapkg/util/slash_util"
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

	from := slash_util.Format(os.Args[1])
	fmt.Println("from", from)

	wxid, needDecrypt := decrypter.GetWxid(from)
	if needDecrypt {
		fmt.Println("wxid", wxid)
		err := decrypter.Decrypt(from, wxid)
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
