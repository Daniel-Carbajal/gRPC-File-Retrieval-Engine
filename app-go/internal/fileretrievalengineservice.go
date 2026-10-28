package internal

import (
	"context"

	pb "csc376/app/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileRetrievalEngineService struct {
	pb.UnimplementedFileRetrievalEngineServer
	store *IndexStore
	// TO-DO keep track of the client IDs
}

func NewFileRetrievalEngineService(store *IndexStore) *FileRetrievalEngineService {
	// TO-DO implement constructor
	return &FileRetrievalEngineService{
		store: store,
	}
}

func (self *FileRetrievalEngineService) Register(ctx context.Context, req *emptypb.Empty) (*pb.RegisterRep, error) {
	// TO-DO generate a client ID
	//       return the client ID as a RegisterRep reply

	return &pb.RegisterRep{}, nil
}

func (self *FileRetrievalEngineService) ComputeIndex(ctx context.Context, req *pb.IndexReq) (*pb.IndexRep, error) {
	// TO-DO extract the document path, client ID and word frequencies from the IndexReq request
	//       get the document number associated with the document path (call PutDocument)
	//       update the index store with the word frequencies and the document number
	//       return an acknowledgement as an IndexRep reply

	return &pb.IndexRep{}, nil
}

func (self *FileRetrievalEngineService) ComputeSearch(ctx context.Context, req *pb.SearchReq) (*pb.SearchRep, error) {
	// TO-DO extract the terms from the SearchReq request
	//       for each term get the pairs of documents and frequencies from the index store
	//       combine the returned documents and frequencies from all of the specified terms
	//       sort the document and frequency pairs and keep only the top 10
	//       for each document number get from the index store the document path
	//       return the top 10 results as a SearchRep reply

	return &pb.SearchRep{}, nil
}

