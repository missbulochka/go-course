# secondHW
## Requirements

- go 1.21.1
- docker 24.0.7

## Build
You can build a project manually with docker. You have to build an image and run a container with commands:
```
docker build \
    -t app-build-base:0.1.0 \
    -f ./secondHW/build.Dockerfile \
    ./secondHW/

docker run \
    --rm \
    -v $(pwd)/secondHW:/go/src/ \
    app-build-base:0.1.0 \
    go run ./cmd/main.go
```

## Run
You can run the application with docker:
```
docker build \
    -t app-build-base:0.1.0 \
    -f ./secondHW/build.Dockerfile \
    ./secondHW/
    
docker build \
    -t app:0.1.0 \
    -f ./secondHW/Dockerfile \
    ./secondHW/

docker run \
    --rm \
    app:0.1.0
```
