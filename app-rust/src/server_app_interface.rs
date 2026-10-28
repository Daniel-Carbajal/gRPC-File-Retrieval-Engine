use std::io::{self, Write};
use std::sync::Arc;

use crate::server_processing_engine::ServerProcessingEngine;

pub struct ServerAppInterface {
    // keep a reference to the server processing engine
    engine: Arc<ServerProcessingEngine>,
}

impl ServerAppInterface {
    pub fn new(engine: Arc<ServerProcessingEngine>) -> Self {
        Self { engine }
    }

    pub fn read_commands(&self) {
        // TO-DO implement the read commands method

        loop {
            print!("> ");
            io::stdout().flush().expect("Failed to flush output");

            // read from command line
            let mut command = String::new();
            let bytes_read = io::stdin().read_line(&mut command);

            if bytes_read.is_err() {
                eprintln!("Unable to read command input.");
                break;
            }

            let command = command.trim();

            // if the command is quit, terminate the program
            if command == "quit" {
                self.engine.shutdown();
                break;
            }

            println!("unrecognized command!");
        }
    }
}

