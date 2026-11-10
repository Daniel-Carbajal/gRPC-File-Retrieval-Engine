<h1>File Retrieval Engine</h1>A client-server file retrieval engine built in Go using gRPC and Protocol Buffers. Clients recursively read files, count term frequencies, and send partial indexes to a central server through gRPC remote procedure calls. The server maintains a shared inverted index and supports searches for documents containing all specified terms. Searches return up to 10 results, ranked by combined term frequency.

Each client uses four indexing worker goroutines and a shared document queue. The gRPC server handles remote procedure calls through the FileRetrievalEngineService implementation. A separate benchmark program measures indexing performance with multiple clients.

Built for a distributed systems course, with a focus on remote procedure calls, Protocol Buffers, concurrent programming, inverted indexes, and performance measurement.

## CSC 376 Programming Assignment 5
**Jarvis College of Computing and Digital Media - DePaul University**

**Student**: dcarbaj4@depaul.edu
**Solution programming language**: Go

<h2>Compile and run instructions:</h2>
Requires Go compatible with the project's go.mod, the Protocol Buffers compiler (protoc), and the Go code-generation plugins. Perform these steps from the repository root.

<h3>1. Enter the Go project directory:</h3>

```bash
cd app-go
```

<h3>2. Install the Protocol Buffers tools:</h3>

On Ubuntu, install protoc if it is not already available:

```bash
sudo apt update
sudo apt install protobuf-compiler
```

Install the Go plugins and add them to the current terminal's PATH:


```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

If needed, add the export line to your shell startup file so the plugins remain available in new terminals. These plugin setup commands follow the <a href="https://grpc.io/docs/languages/go/quickstart/">official gRPC Go quick start.</a>

<h3>3. Generate the Go message types and RPC code:</h3>

From the app-go directory:

```bash
go generate ./proto
go mod download
```

This generates file_retrieval_engine.pb.go and file_retrieval_engine_grpc.pb.go in proto. Regenerate these files after changing file_retrieval_engine.proto; do not edit the generated files directly.

<h3>4. Build the executables:</h3>

```bash
mkdir -p build
go build -o build/file-retrieval-server ./server
go build -o build/file-retrieval-client ./client
go build -o build/file-retrieval-benchmark ./benchmark
```

<h3>5. Start the server:</h3>

From the app-go directory:

```bash
./build/file-retrieval-server <port_number>
```

The argument specifies the port on which the gRPC server listens. Use a port from 1024 through 65535 that is not already in use.

<h3>6. Start a client:</h3>

Open another terminal in the app-go directory:

```bash
./build/file-retrieval-client
```

Connect to the running server from the client prompt:

```Plain text
connect 127.0.0.1 <port_number>
```

Use 127.0.0.1 when the client and server run on the same machine. Otherwise, provide the server's reachable IP address. 

Provide the servers reachable IP address and port number it listens for incoming client connections on. Port must match servers listening port.

<h2>Using the File Retrieval Engine</h2>
<h3>Server commands:</h3>
<ul>
    <li><code>quit</code> — gracefully shut down the gRPC server, allowing active RPCs to finish.</li>
</ul>
<h3>Client commands:</h3>
<ul>
    <li><code>connect &lt;serverIP&gt; &lt;serverPort&gt;</code> — connect to the server and register for a client ID.</li>
    <li><code>get_info</code> — display the assigned client ID.</li>
    <li><code>index &lt;folderPath&gt;</code> — recursively index files in the specified local directory and send their word frequencies to the server.</li>
    <li><code>search &lt;term&gt;</code> — search the server's index for documents containing a term.</li>
    <li><code>search &lt;term1&gt; AND &lt;term2&gt;</code> — search for documents containing all specified terms, with a maximum of three terms.</li>
    <li><code>quit</code> — disconnect from the server and close the client.</li>
</ul>

For example, with the provided datasets stored in the repository's datasets directory and the client started from app-go:

```Plain text
connect 127.0.0.1 8080
get_info
index ../datasets/dataset1_client_server/1_client/client_1
search affidavit
search affidavit AND court
quit
```

Folder paths may be absolute or relative to the directory where the client was started. Files are read on the client machine.

<h2>Running the benchmark</h2>

Start the server first. From the app-go directory, run the benchmark with the server IP, server port, number of clients, and one dataset folder for each client:

```Plain text
./build/file-retrieval-benchmark <serverIP> <serverPort> <numberOfClients> <folder1> ... <folderN>
```

The benchmark accepts one to four clients. The number of folder paths must equal the number of clients.

Example with one client:

```bash
./build/file-retrieval-benchmark 127.0.0.1 8080 1 ../datasets/dataset1_client_server/1_client/client_1
```

Example with two clients:

```bash
./build/file-retrieval-benchmark 127.0.0.1 8080 2 ../datasets/dataset1_client_server/2_clients/client_1 ../datasets/dataset1_client_server/2_clients/client_2
```

Each benchmark worker connects a separate client and indexes its assigned folder. After all workers finish, the benchmark prints the total bytes read, elapsed wall time, and indexing throughput. It then runs three search queries through the first client and disconnects all clients.

Throughput is calculated as:

```Plain text
Throughput (MB/s) = total bytes read / 1,000,000 / elapsed seconds
```

For performance evaluation, run each of the three datasets with 1, 2, and 4 clients. Use the matching dataset partition for each client count. Restart the server before each run to begin with an empty index.

<h2>Indexing and search behavior</h2>
<ul>
    <li>Indexed terms contain four or more characters, consisting of letters, digits, underscores, or hyphens.</li>
    <li>Terms are case-sensitive.</li>
    <li>Multi-term searches use uppercase AND and return only documents containing every specified term.</li>
    <li>Results are ranked by descending combined term frequency, with ties ordered alphabetically by document path.</li>
    <li>Search results identify the originating client and the document's relative path.</li>
    <li>Clients exchange Protocol Buffers messages through gRPC. The current implementation uses an unencrypted channel.</li>
    <li>The server stores the shared index in memory. Restart the server to begin with an empty index.</li>
</ul>

<h2>Remote procedures</h2>

The service and message types are defined in proto/file_retrieval_engine.proto.

| Procedure | Request | Reply |
| --- | --- | --- |
| Register | Empty request | RegisterRep containing the assigned client ID |
| ComputeIndex | IndexReq containing the client ID, relative document path, and word-frequency map | IndexRep containing an acknowledgement |
| ComputeSearch | SearchReq containing the search terms | SearchRep mapping document paths to combined term frequencies |

<h2>Program structure</h2>
<ul>
    <li><strong>Server</strong> — initializes the server components and starts the server command interface.</li>
    <li><strong>ServerAppInterface</strong> — reads the server shutdown command.</li>
    <li><strong>ServerProcessingEngine</strong> — initializes the gRPC server, registers the service, and manages server startup and shutdown.</li>
    <li><strong>FileRetrievalEngineService</strong> — implements registration, indexing, and search RPCs on the server.</li>
    <li><strong>IndexStore</strong> — stores document paths and the term inverted index, protecting shared access with mutexes.</li>
    <li><strong>Client</strong> — initializes the client components and starts the client command interface.</li>
    <li><strong>ClientAppInterface</strong> — parses client commands and displays results.</li>
    <li><strong>ClientProcessingEngine</strong> — manages the gRPC channel and client stub, counts local term frequencies, and calls the server RPCs.</li>
    <li><strong>Proto</strong> — defines the service and messages and contains the generated Go code.</li>
    <li><strong>Benchmark</strong> — runs multiple clients concurrently and reports indexing performance.</li>
</ul>