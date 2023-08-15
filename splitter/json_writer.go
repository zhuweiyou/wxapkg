package splitter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zhuweiyou/wxapkg/formatter"
	"os"
	"path"
	"strings"
)

func WriteAppJson(from string, appConfig gjson.Result) error {
	fmt.Println("write app.json")
	appConfigMap := appConfig.Map()
	appConfigGlobalMap := appConfigMap["global"].Map()
	delete(appConfigMap, "page")
	delete(appConfigMap, "entryPagePath")
	delete(appConfigMap, "debug")
	delete(appConfigMap, "global")
	appJsonMap := make(map[string]any)
	for key, value := range appConfigMap {
		appJsonMap[key] = value.Value()
	}
	for key, value := range appConfigGlobalMap {
		appJsonMap[key] = value.Value()
	}

	appJsonPath := path.Join(from, "app.json")
	err := os.WriteFile(appJsonPath, formatter.FormatJson(appJsonMap), 0666)
	if err != nil {
		return fmt.Errorf("write %s err: %v", appJsonPath, err)
	}

	return nil
}

func WritePageJson(from string, appConfig gjson.Result) error {
	fmt.Println("write page json")
	for pagePath, pageConfig := range appConfig.Get("page").Map() {
		fmt.Println(pagePath, pageConfig)
		err := os.WriteFile(strings.Replace(path.Join(from, pagePath), ".html", ".json", 1), formatter.FormatJsonString(pageConfig.String()), 0666)
		if err != nil {
			return fmt.Errorf("write %s err: %v", pagePath, err)
		}
	}

	return nil
}

func WriteComponentJson(from string, appConfig gjson.Result) error {
	fmt.Println("write component json")
	return nil
}
