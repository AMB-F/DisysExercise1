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
