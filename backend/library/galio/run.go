package galio

func (app *Application) Run() error {
	err := app.runStartHooks()
	if err != nil {
		return err
	}

	return nil
}
