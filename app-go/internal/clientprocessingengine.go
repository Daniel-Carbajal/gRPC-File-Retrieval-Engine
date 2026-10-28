package internal

import (
	pb "csc376/app/proto"
	"google.golang.org/grpc"
)

type IndexResult struct {
	ExecutionTime  float64
	TotalBytesRead int64
}

type DocPathFreqPair struct {
	DocumentPath  string
	WordFrequency int64
}

type SearchResult struct {
	ExecutionTime       float64
	DocumentFrequencies []DocPathFreqPair
}

type ClientProcessingEngine struct {
	// keep track of the connection
	channel *grpc.ClientConn
	stub    pb.FileRetrievalEngineClient
	// TO-DO keep track of the client ID
}

func NewClientProcessingEngine() *ClientProcessingEngine {
	return &ClientProcessingEngine{
		// TO-DO add information to keep track of the gRPC client stub
	}
}

func (self *ClientProcessingEngine) IndexFolder(folderPath string) IndexResult {
	result := IndexResult{ExecutionTime: 0, TotalBytesRead: 0}
	// TO-DO get the start time
	// TO-DO crawl the folder path and extrac all file paths
	// TO-DO for each file extract all words/terms and count their frequencies
	// TO-DO increment the total number of bytes read
	// TO-DO for each file perform a remote procedure call to the server by calling the gRPC client stub
	// TO-DO get the stop time and calculate the execution time
	// TO-DO return the execution time and the total number of bytes read

	return result
}

func (self *ClientProcessingEngine) Search(terms []string) SearchResult {
	result := SearchResult{ExecutionTime: 0, DocumentFrequencies: []DocPathFreqPair{}}
	// TO-DO get the start time
	// TO-DO perform a remote procedure call to the server by calling the gRPC client stub
	// TO-DO get the stop time and calculate the execution time
	// TO-DO return the execution time and the top 10 documents and frequencies

	return result
}

func (self *ClientProcessingEngine) GetInfo() int64 {
	// TO-DO return the client ID

	return 0
}

func (self *ClientProcessingEngine) Connect(serverIP string, serverPort string) {
	// TO-DO create communication channel with the gRPC Server
	// TO-DO create gRPC client stub
	// TO-DO perform a remote procedure call to the server by calling the gRPC client stub
	// TO-DO store the client ID
}

