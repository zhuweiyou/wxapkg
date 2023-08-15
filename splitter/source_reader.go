package splitter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path"
)

func ReadAppConfigJson(from string) (gjson.Result, error) {
	fmt.Println("read app-config.json")
	bytes, err := os.ReadFile(path.Join(from, "app-config.json"))
	if err != nil {
		return gjson.Result{}, fmt.Errorf("read app-config.json err: %v", err)
	}

	return gjson.ParseBytes(bytes), err
}

func ReadAppServiceJs(from string) (string, error) {
	fmt.Println("read app-service.js")
	bytes, err := os.ReadFile(path.Join(from, "app-service.js"))
	if err != nil {
		return "", fmt.Errorf("read app-service.js err: %v", err)
	}

	return string(bytes), err
}
