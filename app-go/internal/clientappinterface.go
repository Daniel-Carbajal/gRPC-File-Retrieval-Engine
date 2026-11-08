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

InputLoop:
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
			if self.engine.channel != nil {
				self.engine.channel.Close()
			}
			break
		}

		// if the command begins with connect, connect to the given server
		if strings.HasPrefix(command, "connect ") {
			tokens := strings.Fields(command)

			if len(tokens) != 3 {
				fmt.Println("Invalid number of arguments; USE: connect <IP Address> <Port Number>")
				continue
			}

			// TO-DO parse command and call Connect on the processing engine
			ipString := tokens[1]
			portString := tokens[2]
			self.engine.Connect(ipString, portString)

			// TO-DO print connection successful
			if self.engine.channel != nil {
				fmt.Printf("Connection successful!\n")
			} else {
				fmt.Println("Connection failed!")
			}

			continue
		}

		// if the command begins with get_info, print the client ID
		if command == "get_info" {
			// TO-DO parse command and call GetInfo on the processing engine
			// TO-DO print the client ID
			fmt.Printf("client ID: %d\n", self.engine.GetInfo())
			continue
		}

		// if the command begins with index, index the files from the specified directory
		if strings.HasPrefix(command, "index ") {
			// TO-DO parse command and call IndexFolder on the processing engine
			tokens := strings.SplitN(command, " ", 2)
			if len(tokens) != 2 {
				fmt.Println("Invalid number of arguments; USE: index <folder_path>")
				continue
			}

			docPath := tokens[1]
			docPath = strings.Trim(docPath, " ")
			result := self.engine.IndexFolder(docPath)

			// TO-DO print the execution time and the total number of bytes read
			fmt.Printf("Completed indexing %d bytes of data\n", result.TotalBytesRead)
			fmt.Printf("Completed indexing in %f seconds\n", result.ExecutionTime)
			continue
		}

		// if the command begins with search, search for files that matches the query
		if strings.HasPrefix(command, "search ") {
			// TO-DO parse command and call Search on the processing engine
			tokens := strings.Fields(command)
			tokens = tokens[1:] // get rid of search

			if tokens[len(tokens)-1] == "AND" && len(tokens)%2 == 0 {
				fmt.Println("Incomplete command: missing final term")
				continue
			}

			var terms []string
			for i, term := range tokens {
				if i%2 == 1 {
					if term != "AND" {
						fmt.Println("Incorrect syntax; unrecognized operation: ", term)
						continue InputLoop
					} else {
						continue
					}
				} else if i%2 == 0 {
					terms = append(terms, term)
				}
			}

			if len(terms) > 3 {
				fmt.Println("Too many terms, maximum 3 allowed; use: 'search <AND query>'")
				continue
			}

			SearchResult := self.engine.Search(terms)

			// TO-DO print the execution time and the top-10 search results
			fmt.Printf("Search execution time is %f \n", SearchResult.ExecutionTime)
			fmt.Printf("Search results for query are (total %v)\n", len(SearchResult.DocumentFrequencies))
			for _, dfp := range SearchResult.DocumentFrequencies {
				fmt.Printf("* %s:%v\n", dfp.DocumentPath, dfp.WordFrequency)
			}

			continue
		}

		fmt.Println("unrecognized command!")
	}

	_ = reader.Err()
}
