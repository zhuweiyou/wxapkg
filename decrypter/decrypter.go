package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"os"
	"strings"
)

type DecryptOptions struct {
	Wxid          string
	Iv            string
	Salt          string
	WxapkgPath    string
	DecWxapkgPath string
}

func Decrypt(options *DecryptOptions) error {
	dataByte, err := os.ReadFile(options.WxapkgPath)
	if err != nil {
		log.Fatal(err)
	}

	dk := pbkdf2.Key([]byte(options.Wxid), []byte(options.Salt), 1000, 32, sha1.New)
	block, _ := aes.NewCipher(dk)
	blockMode := cipher.NewCBCDecrypter(block, []byte(options.Iv))
	originData := make([]byte, 1024)
	blockMode.CryptBlocks(originData, dataByte[6:1024+6])

	afData := make([]byte, len(dataByte)-1024-6)
	var xorKey = byte(0x66)
	if len(options.Wxid) >= 2 {
		xorKey = options.Wxid[len(options.Wxid)-2]
	}
	for i, b := range dataByte[1024+6:] {
		afData[i] = b ^ xorKey
	}

	originData = append(originData[:1023], afData...)

	err = os.WriteFile(options.DecWxapkgPath, originData, 0666)
	if err != nil {
		return fmt.Errorf("write file error: %v", err)
	}

	return nil
}

const DefaultDecryptTo = "_decrypt"

func DefaultDecrypt(from string, wxid string) error {
	return Decrypt(&DecryptOptions{
		Wxid:          wxid,
		Iv:            "the iv: 16 bytes",
		Salt:          "saltiest",
		WxapkgPath:    from,
		DecWxapkgPath: from + DefaultDecryptTo,
	})
}

func GetWxid(from string) (string, bool) {
	parts := strings.Split(from, "/")

	var wxid string
	wxidIndex := len(parts) - 3
	if wxidIndex >= 0 {
		wxid = parts[wxidIndex]
	}

	return wxid, strings.HasPrefix(wxid, "wx")
}
