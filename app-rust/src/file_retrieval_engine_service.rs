use std::sync::Arc;
use tonic::{Request, Response, Status};

use crate::fre::file_retrieval_engine_server::FileRetrievalEngine;
use crate::fre::{IndexRep, IndexReq, RegisterRep, SearchRep, SearchReq};
use crate::index_store::IndexStore;

pub struct FileRetrievalEngineService {
    store: Arc<IndexStore>,
    // TO-DO keep track of the client IDs
}

impl FileRetrievalEngineService {
    pub fn new(store: Arc<IndexStore>) -> Self {
        // TO-DO implement constructor
        Self { store }
    }
}

#[tonic::async_trait]
impl FileRetrievalEngine for FileRetrievalEngineService {
    async fn register(&self, _request: Request<()>) -> Result<Response<RegisterRep>, Status> {
        // TO-DO generate a client ID
        //       return the client ID as a RegisterRep reply

        let reply = RegisterRep { client_id: 0 };
        Ok(Response::new(reply))
    }

    async fn compute_index(
        &self,
        _request: Request<IndexReq>,
    ) -> Result<Response<IndexRep>, Status> {
        // TO-DO extract the document path, client ID and word frequencies from the IndexReq request
        //       get the document number associated with the document path (call put_document)
        //       update the index store with the word frequencies and the document number
        //       return an acknowledgement as an IndexRep reply

        let reply = IndexRep {
            ack: String::new(),
        };
        Ok(Response::new(reply))
    }

    async fn compute_search(
        &self,
        _request: Request<SearchReq>,
    ) -> Result<Response<SearchRep>, Status> {
        // TO-DO extract the terms from the SearchReq request
        //       for each term get the pairs of documents and frequencies from the index store
        //       combine the returned documents and frequencies from all of the specified terms
        //       sort the document and frequency pairs and keep only the top 10
        //       for each document number get from the index store the document path
        //       return the top 10 results as a SearchRep reply

        let reply = SearchRep {
            search_results: std::collections::HashMap::new(),
        };
        Ok(Response::new(reply))
    }
}

