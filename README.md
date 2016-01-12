This's a very simple implementation of grpc <-> rust proxy for "Hello world!" example.
http://www.grpc.io/docs/

It's written for educational purposes.

Usage:
```
$ git clone https://github.com/ikruglov/grpcdemo2rust.git
$ cd grpcdemo2rust
$ cargo build
$ ./src/start.sh

# in parallel bash
$ cd go-proxy
$ ./start

# in another parallel bash
$ cd go-client
$ go run *go -str="Hello world!!!!"
```
