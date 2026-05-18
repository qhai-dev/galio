package main

import (
	"context"
	"fmt"

	"github.com/qhai-dev/galio/library/galio"
)

func main() {
	app := galio.NewApplication()

	app.StartHook(func(ctx context.Context) error {
		err := initializer(app)
		if err != nil {
			return err
		}
		return nil
	})

	app.StartHook(func(ctx context.Context) error {
		fmt.Printf("stathook 2 \n")
		return nil
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}
