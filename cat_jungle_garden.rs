#![feature(decl_macro)]

use std::io::Write;

// A macro to print the current date and time.
macro_rules! date_time {
    () => {
        let now = std::time::SystemTime::now();
        println!("{}", now.format("%Y-%m-%d %H:%M:%S"));
    };
}

// A macro to create a new file with a specified name and content
macro_rules! create_file {
    ($name:expr, $data:expr) => {
        let mut file = std::fs::File::create($name).unwrap();
        file.write_all($data.as_bytes()).unwrap();
    };
}

// A macro to log a message to a log file
macro_rules! log_message {
    ($message:expr) => {
        let log_file = std::fs::File::create("log.txt").unwrap();
        writeln!(&log_file, $message).unwrap();
    };
}

fn main() {
    // Print the current date and time
    date_time!();

    // Create a new file
    create_file!("new_file.txt", "Hello World!");

    // Log a message
    log_message!("Message logged at: ");
}