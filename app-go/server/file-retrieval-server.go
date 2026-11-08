package main

import (
	"csc376/app/internal"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Unexpected number of args")
		return
	}

	// TO-DO change server port to a non-privileged port from os.Args[1]
	serverPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error converting port number arg: ", err)
		return
	}

	if serverPort < 1024 || serverPort > 65535 {
		fmt.Println("Invalid server port number. Valid range: 1024 - 65535")
		return
	}

	store := internal.NewIndexStore()
	engine := internal.NewServerProcessingEngine(store)
	appInterface := internal.NewServerAppInterface(engine)

	// create a thread that runs the gRPC server
	engine.Initialize(serverPort)

	// read commands from the user
	appInterface.ReadCommands()
}
