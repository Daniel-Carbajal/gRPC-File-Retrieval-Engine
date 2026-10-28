package main

import "csc376/app/internal"

func main() {
	engine := internal.NewClientProcessingEngine()
	appInterface := internal.NewClientAppInterface(engine)

	// read commands from the user
	appInterface.ReadCommands()
}

