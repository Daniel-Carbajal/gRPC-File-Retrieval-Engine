package internal

import (
	"cmp"
	"container/list"
	"context"
	pb "csc376/app/proto"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var TERMINATION int64 = -1

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

type bytesRead struct {
	size  int64
	Mutex sync.Mutex
}

type DocQueue struct { // from pa3 for index worker threads
	Queue *list.List
	Mutex sync.Mutex
}

type ClientProcessingEngine struct {
	// keep track of the connection
	channel      *grpc.ClientConn
	stub         pb.FileRetrievalEngineClient
	channelMutex sync.Mutex

	// TO-DO keep track of the client ID
	clientID int32

	// Keep track of index worker threads
	numWorkers int
}

func NewClientProcessingEngine() *ClientProcessingEngine {
	return &ClientProcessingEngine{
		// TO-DO add information to keep track of the gRPC client stub
		numWorkers: 4,
	}
}

func (self *ClientProcessingEngine) IndexFolder(folderPath string) IndexResult {
	result := IndexResult{ExecutionTime: 0, TotalBytesRead: 0}
	size := bytesRead{size: 0}
	var err error

	// TO-DO get the start time
	startTime := time.Now()

	// Create shared Queue of documents for the workers to pop and index from
	filePaths := DocQueue{list.New(), sync.Mutex{}}

	// TO-DO create the worker threads and give to each worker thread a subset of the documents that need to be indexed
	var wg sync.WaitGroup

	if self.channel == nil || self.stub == nil {
		fmt.Println("Connect to a server first")
		return result
	}

	fmt.Println(self.numWorkers, " workers starting . . . ")
	for i := range self.numWorkers {
		wg.Add(1)
		go run_indexing_worker(i+1, self, folderPath, &filePaths, &wg, &size)
		fmt.Println("Started worker number ", i, ". . .")
	}

	// TO-DO crawl the folder path and extrac all file paths
	fmt.Println("Starting folder path crawl. . .")
	err = filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("failure accessing a path %q: %v\n", path, err)
			return err
		}

		if d.Type().IsRegular() {
			filePaths.Mutex.Lock()
			filePaths.Queue.PushBack(path)
			filePaths.Mutex.Unlock()
		}
		return nil
	})

	if err != nil {
		fmt.Println("error walking the path, quiting index and terminating threads:", folderPath, err)
		terminateWorkers(&filePaths, self.numWorkers, true)
		wg.Wait()
		return result
	}

	terminateWorkers(&filePaths, self.numWorkers, false)

	fmt.Println("Folder path crawl finished.")

	// TO-DO join all of the worker threads
	wg.Wait()
	fmt.Println("All workers finished.")

	// TO-DO get the stop time and calculate the execution time
	// TO-DO return the execution time and the total number of bytes read
	result = IndexResult{time.Since(startTime).Seconds(), size.size}
	return result
}

func run_indexing_worker(thread_id int, p *ClientProcessingEngine, folderPath string, docQueue *DocQueue, waitGroup *sync.WaitGroup, size *bytesRead) {
	defer waitGroup.Done()

	// TO-DO for each file extract all words/terms and count their frequencies
	re := regexp.MustCompile(`[a-zA-Z0-9_-]+`)
	var element *list.Element

	for {
		docQueue.Mutex.Lock()
		qLen := docQueue.Queue.Len()

		if qLen == 0 { // Thread can not pop unless the queue has elements in it
			docQueue.Mutex.Unlock()
			continue
		} else {
			// Pop a doc off of the queue
			element = docQueue.Queue.Back()
			docQueue.Queue.Remove(element)
		}
		docQueue.Mutex.Unlock()

		// Check if the element popped is a TERMINATION, if so break loop and end worker
		if element.Value == TERMINATION {
			break
		}

		path := element.Value.(string)

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("failed to read file:", err)
			continue
		}

		relativePath, err := filepath.Rel(folderPath, path)
		if err != nil {
			fmt.Println("failed to resolve relative path:", err)
			continue
		}

		// TO-DO for each file prepare an INDEX REQUEST message and send to the server
		//       the document path, the client ID and the word frequencies
		var indexRequest pb.IndexReq
		indexRequest.ClientId = p.clientID
		indexRequest.DocumentPath = relativePath

		text := string(content) // convert byte slice to string

		rawTerms := re.FindAllString(text, -1) // find all strings that match the regex

		wordFrequencies := make(map[string]int64)
		for _, term := range rawTerms {
			if len(term) > 3 { // match valid sequences of 4 or more characters
				wordFrequencies[term] = wordFrequencies[term] + 1 // increment in the map
			}
		}

		indexRequest.WordFrequencies = wordFrequencies
		if p.channel == nil {
			terminateWorkers(docQueue, p.numWorkers, true)
			return
		}

		// Send index request
		indexReply, err := p.stub.ComputeIndex(context.Background(), &indexRequest)
		if err != nil {
			fmt.Println("Error sending Index Request!", err)
			terminateWorkers(docQueue, p.numWorkers, true)
			return
		}

		// TO-DO receive for each INDEX REQUEST message an INDEX REPLY message
		if indexReply.Ack != "ACK" {
			fmt.Println("Unexpected Server Reply!")
			terminateWorkers(docQueue, p.numWorkers, true)
			return
		}

		// TO-DO increment the total number of bytes read
		size.Mutex.Lock()
		size.size += int64(len(content))
		size.Mutex.Unlock()
	}
}

func terminateWorkers(filePaths *DocQueue, numWorkers int, isNow bool) {
	if isNow {
		for range numWorkers {
			filePaths.Mutex.Lock()
			filePaths.Queue.PushBack(TERMINATION)
			filePaths.Mutex.Unlock()
		}
	} else {
		for range numWorkers {
			filePaths.Mutex.Lock()
			filePaths.Queue.PushFront(TERMINATION)
			filePaths.Mutex.Unlock()
		}
	}
}

func (self *ClientProcessingEngine) Search(terms []string) SearchResult {
	result := SearchResult{ExecutionTime: 0, DocumentFrequencies: []DocPathFreqPair{}}

	// TO-DO get the start time
	startTime := time.Now()

	// TO-DO perform a remote procedure call to the server by calling the gRPC client stub
	var searchRequest pb.SearchReq
	searchRequest.Terms = terms

	// Send search request
	if self.stub == nil {
		fmt.Println("Connect to server first")
		return result
	}

	searchReply, err := self.stub.ComputeSearch(context.Background(), &searchRequest)
	if err != nil {
		fmt.Println("Error sending Search Request!", err)
		return result
	}

	var sortedDocs []DocPathFreqPair

	// Convert return map into slice of DocPathFreqPair
	for docPath, freq := range searchReply.SearchResults {
		sortedDocs = append(sortedDocs, DocPathFreqPair{DocumentPath: docPath, WordFrequency: freq})
	}

	// TO-DO sort the document and frequency pairs and keep only the top 10
	slices.SortFunc(sortedDocs, func(a DocPathFreqPair, b DocPathFreqPair) int {
		if n := cmp.Compare(b.WordFrequency, a.WordFrequency); n != 0 {
			return n
		}

		return cmp.Compare(a.DocumentPath, b.DocumentPath)
	})

	// TO-DO get the stop time and calculate the execution time
	// TO-DO return the execution time and the top 10 documents and frequencies
	result = SearchResult{time.Since(startTime).Seconds(), sortedDocs}
	return result
}

func (self *ClientProcessingEngine) GetInfo() int64 {
	// TO-DO return the client ID
	return int64(self.clientID)
}

func (self *ClientProcessingEngine) Connect(serverIP string, serverPort string) {
	var err error
	// TO-DO create communication channel with the gRPC Server
	target := net.JoinHostPort(serverIP, serverPort)
	self.channel, err = grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not connect to server! ", err)
		self.channel = nil
		self.stub = nil
		return
	}

	// TO-DO create gRPC client stub
	self.stub = pb.NewFileRetrievalEngineClient(self.channel)

	// TO-DO perform a remote procedure call to the server by calling the gRPC client stub
	regReply, err := self.stub.Register(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failure registering to server! ", err)
		self.channel.Close()
		self.channel = nil
		self.stub = nil
		return
	}

	// TO-DO store the client ID
	self.clientID = regReply.GetClientId()
}
