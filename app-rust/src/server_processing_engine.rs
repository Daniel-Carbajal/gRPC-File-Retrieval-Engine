use std::sync::Arc;

use crate::index_store::IndexStore;

pub struct ServerProcessingEngine {
    // keep a reference to the index store
    store: Arc<IndexStore>,
    // TO-DO keep track of the FileRetrievalEngineService object
    // TO-DO keep track of the gRPC Server object
    // TO-DO keep track of the gRPC Server thread
    // TO-DO use Mutexes to implement internal mutability for the data structures when using threads
}

impl ServerProcessingEngine {
    pub fn new(store: Arc<IndexStore>) -> Self {
        // TO-DO initialize FileRetrievalEngineService and gRPC server objects
        Self { store }
    }

    pub fn initialize(self: &Arc<Self>, server_port: i32) {
        // TO-DO create and start the gRPC server thread (clone self and call run_grpc_server)
    }

    pub fn run_grpc_server(self: &Arc<Self>, server_port: i32) {
        // TO-DO build the gRPC server
        // TO-DO register the FileRetrievalEngineService service with the gRPC server
        // TO-DO start the gRPC server
    }

    pub fn shutdown(&self) {
        // TO-DO shutdown the gRPC server
        // TO-DO wait for the gRPC server to shutdown
        // TO-DO join the gRPC server thread
    }
}

