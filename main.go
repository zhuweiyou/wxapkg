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
	fmt.Println("wxid", wxid)

	err := Decrypt(from, wxid)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Unpack(from + DecryptTo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
