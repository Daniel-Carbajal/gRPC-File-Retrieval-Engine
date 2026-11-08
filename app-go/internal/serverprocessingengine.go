package internal

import (
	pb "csc376/app/proto"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ServerProcessingEngine struct {
	// keep a reference to the index store
	store *IndexStore

	// TO-DO keep track of the FileRetrievalEngineService object
	retrievalEngineService *FileRetrievalEngineService

	// TO-DO keep track of the gRPC Server object
	gRPCServer *grpc.Server

	// TO-DO keep track of the gRPC Server thread
	serverWG sync.WaitGroup
}

func NewServerProcessingEngine(store *IndexStore) *ServerProcessingEngine {
	// TO-DO initialize FileRetrievalEngineService and gRPC server objects
	newFileEngine := NewFileRetrievalEngineService(store)
	newServer := grpc.NewServer()
	return &ServerProcessingEngine{store: store, retrievalEngineService: newFileEngine, gRPCServer: newServer}
}

func (self *ServerProcessingEngine) Initialize(serverPort int) {
	// TO-DO create and start the gRPC server thread
	self.serverWG.Add(1)
	go self.RunGRPCServer(serverPort)
}

func (self *ServerProcessingEngine) RunGRPCServer(serverPort int) {
	defer self.serverWG.Done()

	// TO-DO build the gRPC server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		fmt.Println("Error listening on port: ", serverPort, " ", err)
		return
	}

	// TO-DO register the FileRetrievalEngineService service with the gRPC server
	pb.RegisterFileRetrievalEngineServer(self.gRPCServer, self.retrievalEngineService)

	// TO-DO start the gRPC server
	if err := self.gRPCServer.Serve(listener); err != nil {
		return
	}

}

func (self *ServerProcessingEngine) Shutdown() {
	// TO-DO shutdown the gRPC server
	// TO-DO wait for the gRPC server to shutdown
	if self.gRPCServer != nil {
		self.gRPCServer.GracefulStop()
	}

	// TO-DO join the gRPC server thread
	self.serverWG.Wait()
}
