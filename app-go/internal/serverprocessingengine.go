package internal

type ServerProcessingEngine struct {
	// keep a reference to the index store
	store *IndexStore
	// TO-DO keep track of the FileRetrievalEngineService object
	// TO-DO keep track of the gRPC Server object
	// TO-DO keep track of the gRPC Server thread
}

func NewServerProcessingEngine(store *IndexStore) *ServerProcessingEngine {
	// TO-DO initialize FileRetrievalEngineService and gRPC server objects
	return &ServerProcessingEngine{store: store}
}

func (self *ServerProcessingEngine) Initialize(serverPort int) {
	// TO-DO create and start the gRPC server thread
}

func (self *ServerProcessingEngine) RunGRPCServer(serverPort int) {
	// TO-DO build the gRPC server
	// TO-DO register the FileRetrievalEngineService service with the gRPC server
	// TO-DO start the gRPC server
}

func (self *ServerProcessingEngine) Shutdown() {
	// TO-DO shutdown the gRPC server
	// TO-DO wait for the gRPC server to shutdown
	// TO-DO join the gRPC server thread
}

