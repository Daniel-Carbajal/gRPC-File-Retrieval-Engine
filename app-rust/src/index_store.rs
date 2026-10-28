use std::sync::Mutex;

use std::collections::HashMap;

// Data structure that stores a document number and the number of time a word/term appears in the document
pub struct DocFreqPair {
    pub document_number: u64,
    pub word_frequency: u64,
}

pub struct IndexStore {
    // TO-DO declare data structure that keeps track of the DocumentMap
    // TO-DO declare data structures that keeps track of the TermInvertedIndex
    // TO-DO use Mutexes to implement internal mutability for the data structures when using threads
}

impl IndexStore {
    pub fn new() -> Self {
        // TO-DO initialize the DocumentMap and TermInvertedIndex members
        Self {}
    }

    pub fn put_document(&self, document_path: String) -> u64 {
        let mut document_number = 0;
        // TO-DO assign a unique number to the document path and return the number
        // IMPORTANT! you need to make sure that only one thread at a time can access this method

        document_number
    }

    pub fn get_document(&self, document_number: u64) -> String {
        let mut document_path = String::new();
        // TO-DO retrieve the document path that has the given document number

        document_path
    }

    pub fn update_index(&self, document_number: u64, word_frequencies: &HashMap<String, u64>) {
        // TO-DO update the TermInvertedIndex with the word frequencies of the specified document
        // IMPORTANT! you need to make sure that only one thread at a time can access this method
    }

    pub fn lookup_index(&self, term: &str) -> Vec<DocFreqPair> {
        let mut results = vec![];
        // TO-DO return the document and frequency pairs for the specified term

        results
    }
}

