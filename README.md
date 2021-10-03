# DisysExercise1
To run the project, we can recommend using Docker. First, in the `server` folder, build the image with:
```
docker build -t goserver .
```
Then run the project with:
```
docker run -it -p 8080:8080 goserver
```
