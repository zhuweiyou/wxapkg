package json_writer

import (
	"github.com/tidwall/gjson"
)

func Write(from string, appConfig gjson.Result, appService string) error {
	err := WriteApp(from, appConfig)
	if err != nil {
		return err
	}

	err = WritePage(from, appConfig)
	if err != nil {
		return err
	}

	err = WriteComponent(from, appService)
	if err != nil {
		return err
	}

	return nil
}
