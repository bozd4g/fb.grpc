# fb.grpc
gRPC (Remote Procedure Calls) is an open source remote procedure call (RPC) system initially developed at Google in 2015. It uses HTTP/2 for transport, Protocol Buffers as the interface description language, and provides features such as authentication, bidirectional streaming and flow control, blocking or nonblocking bindings, and cancellation and timeouts.

So how we can implement it and run with a web application at the same time?

## Installation

To run a PostgreSql
```sh
$ docker run --name mypostgres -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=fbgrpc -d -p 5432:5432 postgres
``` 
