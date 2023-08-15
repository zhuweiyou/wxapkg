package splitter

import (
	"github.com/zhuweiyou/wxapkg/formatter"
	"testing"
)

func TestSplit(t *testing.T) {
	from := formatter.FormatSlash("D:\\wechat\\WeChat Files\\Applet\\wxc2ee296e3b2fe7a6\\8\\__APP__.wxapkg_decrypt_unpack")
	t.Log(Split(from))
}