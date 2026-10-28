pub struct IndexResult {
    pub execution_time: f64,
    pub total_bytes_read: u64,
}

pub struct DocPathFreqPair {
    pub document_path: String,
    pub word_frequency: u64,
}

pub struct SearchResult {
    pub execution_time: f64,
    pub document_frequencies: Vec<DocPathFreqPair>,
}

pub struct ClientProcessingEngine {
    // TO-DO keep track of the gRPC client stub and runtime
    // TO-DO keep track of the client ID
    // TO-DO use Mutexes to implement internal mutability when using shared references
}

impl ClientProcessingEngine {
    pub fn new() -> Self {
        // TO-DO add information to keep track of the gRPC client stub
        Self {}
    }

    pub fn index_folder(&self, folder_path: &str) -> IndexResult {
        let mut result = IndexResult {
            execution_time: 0.0,
            total_bytes_read: 0,
        };
        // TO-DO get the start time
        // TO-DO crawl the folder path and extrac all file paths
        // TO-DO for each file extract all words/terms and count their frequencies
        // TO-DO increment the total number of bytes read
        // TO-DO for each file perform a remote procedure call to the server by calling the gRPC client stub
        // TO-DO get the stop time and calculate the execution time
        // TO-DO return the execution time and the total number of bytes read

        result
    }

    pub fn search(&self, terms: Vec<String>) -> SearchResult {
        let mut result = SearchResult {
            execution_time: 0.0,
            document_frequencies: vec![],
        };
        // TO-DO get the start time
        // TO-DO perform a remote procedure call to the server by calling the gRPC client stub
        // TO-DO get the stop time and calculate the execution time
        // TO-DO return the execution time and the top 10 documents and frequencies

        result
    }

    pub fn get_info(&self) -> u64 {
        // TO-DO return the client ID

        0
    }

    pub fn connect(&self, server_ip: &str, server_port: &str) {
        // TO-DO create communication channel with the gRPC Server
        // TO-DO create gRPC client stub
        // TO-DO perform a remote procedure call to the server by calling the gRPC client stub
        // TO-DO store the client ID
    }
}

