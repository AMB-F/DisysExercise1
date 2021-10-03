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
