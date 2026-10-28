package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientAppInterface struct {
	// keep a reference to the client processing engine
	engine *ClientProcessingEngine
}

func NewClientAppInterface(engine *ClientProcessingEngine) *ClientAppInterface {
	return &ClientAppInterface{engine: engine}
}

func (self *ClientAppInterface) ReadCommands() {
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
			// close the program
			break
		}

		// if the command begins with connect, connect to the given server
		if strings.HasPrefix(command, "connect") {
			// TO-DO parse command and call Connect on the processing engine
			// TO-DO print connection successful
			continue
		}

		// if the command begins with get_info, print the client ID
		if strings.HasPrefix(command, "get_info") {
			// TO-DO parse command and call GetInfo on the processing engine
			// TO-DO print the client ID
			continue
		}

		// if the command begins with index, index the files from the specified directory
		if strings.HasPrefix(command, "index") {
			// TO-DO parse command and call IndexFolder on the processing engine
			// TO-DO print the execution time and the total number of bytes read
			continue
		}

		// if the command begins with search, search for files that matches the query
		if strings.HasPrefix(command, "search") {
			// TO-DO parse command and call Search on the processing engine
			// TO-DO print the execution time and the top-10 search results
			continue
		}

		fmt.Println("unrecognized command!")
	}

	_ = reader.Err()
}

