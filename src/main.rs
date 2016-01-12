extern crate hyper;
extern crate protobuf; // depend on rust-protobuf runtime

use std::io::Read;

use hyper::Server;
use hyper::server::Request;
use hyper::server::Response;

mod helloworld;
use protobuf::Message;
use helloworld::HelloReply;
use helloworld::HelloRequest;

fn hello(mut req: Request, res: Response) {
    let mut inb = Vec::new();
    req.read_to_end(&mut inb).unwrap();

    let mut hello_request = HelloRequest::new();
    hello_request.merge_from_bytes(&inb).unwrap();
    let name = hello_request.take_name();
    println!("new request {}", name);

    let mut hello_reply = HelloReply::new();
    hello_reply.set_message(name);
    let outb = hello_reply.write_to_bytes().unwrap();
    res.send(&outb).unwrap();
}

fn main() {
    println!("start hyper http service on 127.0.0.1:3000");
    Server::http("127.0.0.1:3000").unwrap().handle(hello).unwrap();
}
