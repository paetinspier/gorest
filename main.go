package main

import "psn/gorest/app"

func main() {
	// setup and run service

	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
