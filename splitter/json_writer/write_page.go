package json_writer

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zhuweiyou/wxapkg/util/json_util"
	"os"
	"path"
	"strings"
)

func WritePage(from string, appConfig gjson.Result) error {
	fmt.Println("write page json")
	for pagePath, pageConfig := range appConfig.Get("page").Map() {
		fmt.Println(pagePath, pageConfig)
		err := os.WriteFile(strings.Replace(path.Join(from, pagePath), ".html", ".json", 1), json_util.FormatFromString(pageConfig.String()), 0666)
		if err != nil {
			return fmt.Errorf("write %s err: %v", pagePath, err)
		}
	}

	return nil
}
