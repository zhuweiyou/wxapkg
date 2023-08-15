package splitter

import (
	"github.com/zhuweiyou/wxapkg/splitter/json_writer"
	"github.com/zhuweiyou/wxapkg/splitter/source_reader"
)

func Split(from string) error {
	appConfig, err := source_reader.ReadAppConfig(from)
	if err != nil {
		return err
	}

	appService, err := source_reader.ReadAppService(from)
	if err != nil {
		return err
	}

	err = json_writer.Write(from, appConfig, appService)
	if err != nil {
		return err
	}

	return nil
}
