package csc376.app;

import io.grpc.stub.StreamObserver;

public class FileRetrievalEngineService extends FileRetrievalEngineGrpc.FileRetrievalEngineImplBase {
    private IndexStore store;
    // TO-DO keep track of the client IDs
    
    public FileRetrievalEngineService(IndexStore store) {
        this.store = store;
    }

    @Override
    public void register(com.google.protobuf.Empty request, StreamObserver<RegisterRep> responseObserver) {
        responseObserver.onNext(doRegister());
        responseObserver.onCompleted();
    }

    @Override
    public void computeIndex(IndexReq request, StreamObserver<IndexRep> responseObserver) {
        responseObserver.onNext(doIndex(request));
        responseObserver.onCompleted();
    }

    @Override
    public void computeSearch(SearchReq request, StreamObserver<SearchRep> responseObserver) {
        responseObserver.onNext(doSearch(request));
        responseObserver.onCompleted();
    }

    private RegisterRep doRegister() {
        // TO-DO generate a client ID
        //       return the client ID as a RegisterRep reply

        return RegisterRep.newBuilder().build();
    }

    private IndexRep doIndex(IndexReq request) {
        // TO-DO update global index with temporary index received from the request
        // TO-DO send an OK message as the reply

        return IndexRep.newBuilder().build();
    }

    private SearchRep doSearch(SearchReq request) {
        // TO-DO do lookup over the global index given the search term from the request
        // TO-DO send the results as the reply message

        return SearchRep.newBuilder().build();
    }

}
