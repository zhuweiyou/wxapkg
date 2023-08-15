package splitter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path"
	"strings"
)

func Split(from string) error {
	appConfigBytes, err := os.ReadFile(path.Join(from, "app-config.json"))
	if err != nil {
		return fmt.Errorf("read app-config.json err: %v", err)
	}

	//appConfigList := gjson.GetManyBytes(appConfigBytes, "page", "pages", "global")
	for pagePath, pageConfig := range gjson.GetBytes(appConfigBytes, "page").Map() {
		fmt.Println(pagePath, pageConfig)
		err := os.WriteFile(strings.Replace(path.Join(from, pagePath), ".html", ".json", 1), FormatJson([]byte(pageConfig.String())), 0666)
		if err != nil {
			return fmt.Errorf("write %s err: %v", pagePath, err)
		}
	}

	return nil
}
