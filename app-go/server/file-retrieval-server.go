package main

import "csc376/app/internal"

func main() {
	// TO-DO change server port to a non-privileged port from os.Args[1]
	var serverPort int = 1

	store := internal.NewIndexStore()
	engine := internal.NewServerProcessingEngine(store)
	appInterface := internal.NewServerAppInterface(engine)

	// create a thread that runs the gRPC server
	engine.Initialize(serverPort)

	// read commands from the user
	appInterface.ReadCommands()
}

