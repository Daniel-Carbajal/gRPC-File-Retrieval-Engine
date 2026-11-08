package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ServerAppInterface struct {
	// keep a reference to the server processing engine
	engine *ServerProcessingEngine
}

func NewServerAppInterface(engine *ServerProcessingEngine) *ServerAppInterface {
	return &ServerAppInterface{engine: engine}
}

func (self *ServerAppInterface) ReadCommands() {
	// TO-DO implement the read commands method
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		// read from command line
		if !reader.Scan() {
			break
		}

		command := strings.TrimSpace(reader.Text())

		// if the command is quit, terminate the program
		if command == "quit" {
			self.engine.Shutdown()
			break
		}

		fmt.Println("unrecognized command!")
	}

	_ = reader.Err()
}
