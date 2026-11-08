package internal

import (
	"cmp"
	"context"
	"slices"
	"strconv"
	"sync"

	pb "csc376/app/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type FileRetrievalEngineService struct {
	pb.UnimplementedFileRetrievalEngineServer
	store *IndexStore

	// TO-DO keep track of the client IDs
	clientCounter int32
	counterMutex  sync.Mutex
}

func NewFileRetrievalEngineService(store *IndexStore) *FileRetrievalEngineService {
	// TO-DO implement constructor
	return &FileRetrievalEngineService{
		store:         store,
		clientCounter: 1,
	}
}

func (self *FileRetrievalEngineService) Register(ctx context.Context, req *emptypb.Empty) (*pb.RegisterRep, error) {
	// TO-DO generate a client ID
	var cID int32
	self.counterMutex.Lock()
	cID = self.clientCounter
	self.clientCounter = self.clientCounter + 1
	self.counterMutex.Unlock()

	//       return the client ID as a RegisterRep reply
	return &pb.RegisterRep{ClientId: cID}, nil
}

func (self *FileRetrievalEngineService) ComputeIndex(ctx context.Context, req *pb.IndexReq) (*pb.IndexRep, error) {
	// TO-DO extract the document path, client ID and word frequencies from the IndexReq request
	cID := strconv.Itoa(int(req.ClientId))
	wFreqs := req.WordFrequencies
	docPath := "client" + cID + ":" + req.DocumentPath

	//       get the document number associated with the document path (call PutDocument)
	docNum := self.store.PutDocument(docPath)

	//       update the index store with the word frequencies and the document number
	self.store.UpdateIndex(docNum, wFreqs)

	//       return an acknowledgement as an IndexRep reply
	return &pb.IndexRep{Ack: "ACK"}, nil
}

func (self *FileRetrievalEngineService) ComputeSearch(ctx context.Context, req *pb.SearchReq) (*pb.SearchRep, error) {
	// TO-DO extract the terms from the SearchReq request
	terms := req.Terms

	//       for each term get the pairs of documents and frequencies from the index store
	var prevCandidates []DocFreqPair
	var currCandidates []DocFreqPair
	for i, term := range terms {
		if i == 0 {
			currCandidates = self.store.LookupIndex(term)
			prevCandidates = currCandidates
		} else {
			//       combine the returned documents and frequencies from all of the specified terms
			currCandidates = Intersect(prevCandidates, self.store.LookupIndex(term))
			prevCandidates = currCandidates
		}
	}

	//       for each document number get from the index store the document path
	var sortedDocs []DocPathFreqPair
	for _, val := range currCandidates {
		temp := DocPathFreqPair{self.store.GetDocument(val.DocumentNumber), val.WordFrequency}

		sortedDocs = append(sortedDocs, temp)
	}

	//       sort the document and frequency pairs and keep only the top 10
	slices.SortFunc(sortedDocs, func(a DocPathFreqPair, b DocPathFreqPair) int {
		if n := cmp.Compare(b.WordFrequency, a.WordFrequency); n != 0 {
			return n
		}

		return cmp.Compare(a.DocumentPath, b.DocumentPath)
	})

	if len(sortedDocs) > 10 {
		sortedDocs = sortedDocs[:10]
	}

	reply := make(map[string]int64)
	for _, pair := range sortedDocs {
		reply[pair.DocumentPath] = pair.WordFrequency
	}

	//       return the top 10 results as a SearchRep reply
	return &pb.SearchRep{SearchResults: reply}, nil
}

// Helper function where the returned list of DocFreqPair is an intersection of the two arguments
func Intersect(slice1 []DocFreqPair, slice2 []DocFreqPair) []DocFreqPair {
	var result []DocFreqPair

	tempMap := make(map[int64]int64)
	for _, pair := range slice1 {
		tempMap[pair.DocumentNumber] = pair.WordFrequency
	}

	for _, pair := range slice2 {
		if freq, ok := tempMap[pair.DocumentNumber]; ok {
			temp := DocFreqPair{pair.DocumentNumber, freq + pair.WordFrequency}
			// add to result
			result = append(result, temp)
		}
	}

	return result
}
