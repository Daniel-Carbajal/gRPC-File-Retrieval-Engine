pub mod fre {
    tonic::include_proto!("fre");
}

mod client_app_interface;
mod client_processing_engine;

use std::sync::Arc;

use client_app_interface::ClientAppInterface;
use client_processing_engine::ClientProcessingEngine;

fn main() {
    let engine = Arc::new(ClientProcessingEngine::new());
    let engine_clone = Arc::clone(&engine);
    let interface = ClientAppInterface::new(engine_clone);

    // read commands from the user
    interface.read_commands();
}

