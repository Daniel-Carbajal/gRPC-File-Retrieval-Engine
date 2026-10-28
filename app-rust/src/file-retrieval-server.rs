pub mod fre {
    tonic::include_proto!("fre");
}

mod file_retrieval_engine_service;
mod index_store;
mod server_app_interface;
mod server_processing_engine;

use std::sync::Arc;

use index_store::IndexStore;
use server_app_interface::ServerAppInterface;
use server_processing_engine::ServerProcessingEngine;

fn main() {
    // TO-DO change server port to a non-privileged port from env::args()
    let mut server_port = 1;

    let store = Arc::new(IndexStore::new());
    let store_clone = Arc::clone(&store);
    let engine = Arc::new(ServerProcessingEngine::new(store_clone));
    let engine_clone = Arc::clone(&engine);
    let interface = ServerAppInterface::new(engine_clone);

    // create a thread that runs the gRPC server
    engine.initialize(server_port);

    // read commands from the user
    interface.read_commands();
}

