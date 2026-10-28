use std::io::{self, Write};
use std::sync::Arc;

use crate::client_processing_engine::ClientProcessingEngine;

pub struct ClientAppInterface {
    // keep a reference to the client processing engine
    engine: Arc<ClientProcessingEngine>,
}

impl ClientAppInterface {
    pub fn new(engine: Arc<ClientProcessingEngine>) -> Self {
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
                // close the program
                break;
            }

            // if the command begins with connect, connect to the given server
            if command.starts_with("connect") {
                // TO-DO parse command and call connect on the processing engine
                // TO-DO print connection successful
                continue;
            }

            // if the command begins with get_info, print the client ID
            if command.starts_with("get_info") {
                // TO-DO parse command and call get_info on the processing engine
                // TO-DO print the client ID
                continue;
            }

            // if the command begins with index, index the files from the specified directory
            if command.starts_with("index") {
                // TO-DO parse command and call index_folder on the processing engine
                // TO-DO print the execution time and the total number of bytes read
                continue;
            }

            // if the command begins with search, search for files that matches the query
            if command.starts_with("search") {
                // TO-DO parse command and call search on the processing engine
                // TO-DO print the execution time and the top 10 search results
                continue;
            }

            println!("unrecognized command!");
        }
    }
}

