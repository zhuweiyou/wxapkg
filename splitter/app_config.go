package splitter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path"
)

func ReadAppConfig(from string) (gjson.Result, error) {
	fmt.Println("read app-config.json")
	bytes, err := os.ReadFile(path.Join(from, "app-config.json"))
	if err != nil {
		return gjson.Result{}, fmt.Errorf("read app-config.json err: %v", err)
	}

	return gjson.ParseBytes(bytes), err
}
