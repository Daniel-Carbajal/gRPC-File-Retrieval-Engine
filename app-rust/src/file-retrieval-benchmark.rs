pub mod fre {
    tonic::include_proto!("fre");
}

mod client_app_interface;
mod client_processing_engine;

fn run_worker() {
    // TO-DO connect the ClientProcessingEngine to the server
    // TO-DO index the dataset
}

fn main() {
    let mut server_ip = String::new();
    let mut server_port = String::new();
    let mut number_of_clients = 0;
    let mut clients_dataset_path: Vec<String> = vec![];

    // TO-DO extract the arguments from env::args()
    // TO-DO measure the execution start time
    // TO-DO create ClientProcessingEngine objects equal to the number of clients
    // TO-DO create and start benchmark worker threads equal to the number of clients
    // TO-DO join the benchmark worker threads
    // TO-DO measure the execution stop time and print the performance
    // TO-DO run search queries on the first client (benchmark worker thread number 1)
    // TO-DO disconnect all clients (all benchmakr worker threads)
}

