# DisysExercise1
## Server
To run the server, we can recommend using Docker. First, in the `server` folder, build the image with:
```
docker build -t goserver .
```
Then run the project with:
```
docker run -it -p 8080:8080 goserver
```

Note: the server does not contain actual business logic - i.e. the requests don't actually do anything. It's just a constant response.

## Client
To run the client, enter the `client` folder and run `go run main.go`.



## Protocol Buffers

In the /gRPC/server folder, run:

### Installation

```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Build IDL for the server

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto
```