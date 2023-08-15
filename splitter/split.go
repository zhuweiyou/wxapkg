package splitter

func Split(from string) error {
	appConfig, err := ReadAppConfig(from)
	if err != nil {
		return err
	}

	err = WriteAppJson(from, appConfig)
	if err != nil {
		return err
	}

	err = WritePageJson(from, appConfig)
	if err != nil {
		return err
	}

	return nil
}
