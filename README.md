# DisysExercise1
## Group members
Group name: *Point of Go return*

- Adrian Borup (adbo@itu.dk)
- Andreas Wachs (ahja@itu.dk)
- Anne M. Bartholdy-Falk (anmb@itu.dk)
- Joachim Borup (aljb@itu.dk)

## Server
To run the server, enter the `server` folder and run `go run .`. The server will then be listening at `localhost:8080`.

Note: the server does not contain actual business logic - i.e. the requests don't actually do anything. It's just a constant response.

## Client
To run the client, enter the `client` folder and run `go run main.go`.

## Protocol Buffers

In the /gRPC/server folder, run:

## Example output

### Server
```
2021/10/03 21:47:39 Server started
2021/10/03 21:47:46 GET /course/1 CourseCourseIdGet 0s
2021/10/03 21:47:46 POST /course CoursePost 0s
2021/10/03 21:47:46 PUT /course/1 CourseCourseIdPut 0s
2021/10/03 21:47:46 DELETE /course/1 CourseCourseIdDelete 0s
```

### Client
```
Requesting [GET] http://127.0.0.1:8080/course/1
{
        "id": 0,
        "name": "string",
        "expectedWorkload": 0,
        "rating": 0
}

Requesting [POST] http://127.0.0.1:8080/course
{
        "id": 0,
        "name": "string",
        "expectedWorkload": 0,
        "rating": 0
}

Requesting [PUT] http://127.0.0.1:8080/course/1
{
        "id": 0,
        "name": "string",
        "expectedWorkload": 0,
        "rating": 0
}

Requesting [DELETE] http://127.0.0.1:8080/course/1
{
        "id": 0,
        "name": "string",
        "expectedWorkload": 0,
        "rating": 0
}
```

### Installation

```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Build IDL for the server

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto
```

## Running gRPC with Docker
### Build the image
Build the docker image inside the `gRPC` folder with
```
docker build -t gogrpc .
```

### Run the container
To spin up the container, run:
```bash
# unix
docker run -it --rm -v "$(pwd)":/app --network="host" gogrpc bash
# windows
docker run -it --rm -v ${pwd}:/app --network="host" gogrpc bash
```
This assumes that you are in the `gRPC` folder! Note, if your path has spaces, it may very well not work.

### Compile Protobuf
```
cd /app && \
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional service/service.proto
```

### Run the server
```
cd /app/server && \
  go run .
```

### Run the client
```
cd /app/client && \
  go run .
```
