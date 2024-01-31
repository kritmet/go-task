# Task Management API

Task Management API is a golang project for manage tasks.

## How to run ?

API runnable via Docker. Here's a example:

1. build the docker
```bash
docker build -t go-task .
```
2. Run the Docker container
```bash
docker run -p 8000:8000 go-task
```
OR simple choice:
```bash
go run .
```

The API will be accessible at http://localhost:8000

## API Ducument


The Swagger API documentation is available at 
```
http://localhost:8000/swagger/index.html
```
You can explore and test the API endpoints using the Swagger UI.

