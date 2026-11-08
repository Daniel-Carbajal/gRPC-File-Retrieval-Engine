package main

import (
	"csc376/app/internal"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func runWorker(e *internal.ClientProcessingEngine, serverIP string, serverPort string, datasetPath string, bytesRead *int64, bytesMutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// TO-DO connect the ClientProcessingEngine to the server
	e.Connect(serverIP, serverPort)

	// TO-DO index the dataset
	bytes := e.IndexFolder(datasetPath).TotalBytesRead

	bytesMutex.Lock()
	*bytesRead = *bytesRead + bytes
	bytesMutex.Unlock()
}

func main() {
	var serverIP string
	var serverPort string
	var numberOfClients int
	var clientsDatasetPath []string

	_, _, _, _ = serverIP, serverPort, numberOfClients, clientsDatasetPath

	// TO-DO extract the arguments from os.Args
	userArgs := os.Args[1:]
	if len(userArgs) < 4 || len(userArgs) > 7 {
		fmt.Println("Invalid arguments. Try: <server IP> <server port> <number of clients> <folder 1> ... <folder N (MAX N = 4)>")
		return
	}

	numberOfClients, err := strconv.Atoi(userArgs[2])
	if err != nil {
		fmt.Printf("Invalid number of clients %s: Must be a valid integer 1-4", userArgs[2])
		return
	}
	if numberOfClients < 1 || numberOfClients > 4 {
		fmt.Printf("Invalid number of clients %s: Must be a valid integer 1-4", userArgs[2])
		return
	}

	clientsDatasetPath = userArgs[3:]
	if len(clientsDatasetPath) != numberOfClients {
		fmt.Println("Number of folderpaths provided must equal number of clients")
		return
	}

	serverIP = userArgs[0]
	serverPort = userArgs[1]

	// TO-DO measure the execution start time
	startTime := time.Now()

	// TO-DO create ClientProcessingEngine objects equal to the number of clients
	engines := make([]*internal.ClientProcessingEngine, numberOfClients)

	// populate
	for i := range engines {
		engines[i] = internal.NewClientProcessingEngine()
	}

	// TO-DO create and start benchmark worker threads equal to the number of clients
	var benchmarkWG sync.WaitGroup
	var bytesRead int64
	var bytesMutex sync.Mutex

	for i := range engines {
		benchmarkWG.Add(1)
		go runWorker(engines[i], serverIP, serverPort, clientsDatasetPath[i], &bytesRead, &bytesMutex, &benchmarkWG)
	}

	// TO-DO join the benchmark worker threads
	benchmarkWG.Wait()
	// TO-DO measure the execution stop time and print the performance
	totalTime := time.Since(startTime).Seconds()
	fmt.Printf("Completed indexing %d bytes of data\n", bytesRead)
	fmt.Printf("Completed indexing in %f seconds\n", totalTime)

	throughput := float64(bytesRead) / float64(1000000) / totalTime
	fmt.Printf("Throughput (MB/s): %f\n", throughput)

	// TO-DO run search queries on the first client (benchmark worker thread number 1)
	var terms []string
	var result internal.SearchResult

	terms = append(terms, "affidavit")
	result = engines[0].Search(terms)
	printResults(result)

	terms = append(terms, "eyelashes")
	result = engines[0].Search(terms)
	printResults(result)

	terms = append(terms, "Likewise")
	result = engines[0].Search(terms)
	printResults(result)

	// TO-DO disconnect all clients (all benchmakr worker threads)
	for i := range engines {
		engines[i].Disconnect()
	}
}

func printResults(result internal.SearchResult) {
	fmt.Printf("Search execution time is %f \n", result.ExecutionTime)
	fmt.Printf("Search results for query are (total %v)\n", len(result.DocumentFrequencies))
	for _, dfp := range result.DocumentFrequencies {
		fmt.Printf("* %s:%v\n", dfp.DocumentPath, dfp.WordFrequency)
	}
}
