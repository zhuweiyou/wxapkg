package source_reader

import (
	"fmt"
	"os"
	"path"
)

func ReadAppService(from string) (string, error) {
	fmt.Println("read app-service.js")
	bytes, err := os.ReadFile(path.Join(from, "app-service.js"))
	if err != nil {
		return "", fmt.Errorf("read app-service.js err: %v", err)
	}

	return string(bytes), err
}
