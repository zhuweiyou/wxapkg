package splitter

import (
	"github.com/zhuweiyou/wxapkg/splitter/json_writer"
	"github.com/zhuweiyou/wxapkg/splitter/source_reader"
)

func Split(from string) error {
	appConfig, err := source_reader.ReadAppConfigJson(from)
	if err != nil {
		return err
	}

	err = json_writer.Write(from, appConfig)
	if err != nil {
		return err
	}

	return nil
}
