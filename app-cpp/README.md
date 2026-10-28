### Requirements

To implement your solution in C++ you will need to have GCC 15, CMake, Protobuf and gRPC installed on your system.
On Ubuntu 26.04 LTS you can install GCC, CMake, Protobuf and gRPC using the following commands:

```bash
sudo apt install build-essential cmake g++ gcc pkg-config protobuf-compiler libprotobuf-dev protobuf-compiler-grpc libgrpc++-dev
```

### Setup

There are 3 datasets (dataset1_client_server, dataset2_client_server, dataset3_client_server) that you need to use to evaluate the indexing performance of your solution.
Before you can evaluate your solution you need to download the datasets.
You can download the datasets from the following link:

https://depauledu-my.sharepoint.com/:f:/g/personal/aorhean_depaul_edu/Ej4obLnAKMdFh1Hidzd1t1oBHY7IvgqXoLdKRg-buoiisw?e=SWLALa

After you finished downloading the datasets copy them to the dataset directory (create the directory if it does not exist).
Here is an example on how you can copy Dataset1 to the remote machine and how to unzip the dataset:

```bash
remote-computer$ mkdir datasets
local-computer$ scp dataset1_client_server.zip cc@<remote-ip>:<path-to-repo>/datasets/.
remote-computer$ cd <path-to-repo>/datasets
remote-computer$ unzip dataset1_client_server.zip
```

### How to build/compile

To build the C++ solution use the following commands:

```bash
cd app-cpp
mkdir build
cmake -S . -B build
cmake --build build --config Release
```

### How to run application

To run the C++ server (after you build the project) use the following command:

```bash
./build/file-retrieval-server <port>
> <quit>
```

To run the C++ client (after you build the project) use the following command:

```bash
./build/file-retrieval-client
> <connect | get_info | index | search | quit>
```

To run the C++ benchmark (after you build the project) use the following command:

```bash
./build/file-retrieval-benchmark <server IP> <server port> <number of clients> [<dataset path>]
```

### Example (2 clients and 1 server)

**Step 1:** start the server:

Server

```bash
./build/file-retrieval-server 12345
>
```

**Step 2:** start the clients and connect them to the server:

Client 1

```bash
./build/file-retrieval-client
> connect 127.0.0.1 12345
Connection successful!
> get_info
client ID: 1
>
```

Client 2

```bash
./build/file-retrieval-client
> connect 127.0.0.1 12345
Connection successful!
> get_info
client ID: 2
>
```

**Step 3:** index files from the clients:

Client 1

```bash
> index ../datasets/dataset1_client_server/2_clients/client_1
Completed indexing 68383239 bytes of data
Completed indexing in 2.974 seconds
>
```

Client 2

```bash
> index ../datasets/dataset1_client_server/2_clients/client_2
Completed indexing 65864138 bytes of data
Completed indexing in 2.386 seconds
>
```

**Step 4:** search files from the clients:

Client 1

```bash
> search the
Search completed in 0.4 seconds
Search results (top 10 out of 0):
> search reconcilable
Search completed in 2.1 seconds
Search results (top 10 out of 5):
* client1:client_1/folder4/Document10633.txt:4
* client2:client_2/folder6/Document10807.txt:2
* client2:client_2/folder4/Document10616.txt:1
* client2:client_2/folder4/Document10681.txt:1
* client2:client_2/folder7/folderD/Document11034.txt:1
>
```

Client 2

```bash
> search war-whoops
Search completed in 2.8 seconds
Search results (top 10 out of 5):
* client1:client_1/folder1/Document10086.txt:3
* client1:client_1/folder2/Document10268.txt:1
* client1:client_1/folder2/folderA/Document10370.txt:1
* client1:client_1/folder3/folderB/Document10479.txt:1
* client2:client_2/folder8/Document11154.txt:1
> search amazing AND sixty-eight
Search completed in 3.8 seconds
Search results (top 10 out of 4):
* client2:client_2/folder6/Document10818.txt:14
* client1:client_1/folder4/Document10600.txt:6
* client1:client_1/folder1/Document10060.txt:2
* client1:client_1/folder2/folderA/Document10350.txt:2
>
```

**Step 5:** close and disconnect the clients:

Client 1

```bash
> quit
```

Client 2

```bash
> quit
```

**Step 6:** close the server:

Server

```bash
> quit
```

### Example (benchmark with 2 clients and 1 server)

**Step 1:** start the server:

Server

```bash
./build/file-retrieval-server 12345
>
```

**Step 2:** start the benchmark:

Benchmark

```bash
./build/file-retrieval-benchmark 127.0.0.1 12345 2 ../datasets/dataset1_client_server/2_clients/client_1 ../datasets/dataset1_client_server/2_clients/client_2
Completed indexing 134247377 bytes of data
Completed indexing in 6.015 seconds
Searching the
Search completed in 0.4 seconds
Search results (top 10 out of 0):
Searching search reconcilable
Search completed in 2.1 seconds
Search results (top 10 out of 5):
* client1:client_1/folder4/Document10633.txt:4
* client2:client_2/folder6/Document10807.txt:2
* client2:client_2/folder4/Document10616.txt:1
* client2:client_2/folder4/Document10681.txt:1
* client2:client_2/folder7/folderD/Document11034.txt:1
Searching search war-whoops
Search completed in 2.8 seconds
Search results (top 10 out of 5):
* client1:client_1/folder1/Document10086.txt:3
* client1:client_1/folder2/Document10268.txt:1
* client1:client_1/folder2/folderA/Document10370.txt:1
* client1:client_1/folder3/folderB/Document10479.txt:1
* client2:client_2/folder8/Document11154.txt:1
Searching search amazing AND sixty-eight
Search completed in 3.8 seconds
Search results (top 10 out of 4):
* client2:client_2/folder6/Document10818.txt:14
* client1:client_1/folder4/Document10600.txt:6
* client1:client_1/folder1/Document10060.txt:2
* client1:client_1/folder2/folderA/Document10350.txt:2
```

**Step 3:** close the server:

Server

```bash
> quit
```

### Autograding

Install Python 3.14:

```bash
sudo apt install python3.14
```

Run the autograder driver:

```bash
python3.14 test/setup.py
python3.14 test/test_driver.py
```
