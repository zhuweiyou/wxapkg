package splitter

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path"
)

func ReadAppServiceJs(from string) (gjson.Result, error) {
	fmt.Println("read app-service.js")
	bytes, err := os.ReadFile(path.Join(from, "app-service.js"))
	if err != nil {
		return gjson.Result{}, fmt.Errorf("read app-service.js err: %v", err)
	}

	return gjson.ParseBytes(bytes), err
}
