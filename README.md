# containerization-steps
The steps to containerize an application

## Dockerfile explanation

```Dockerfile
# use golang image as the base image
FROM golang 

# specify the working directory
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

# copy all the source file into the image
COPY . .

# compile the app
RUN go build -v -o /usr/local/bin/app ./...

# start the app
CMD ["app"]

# serve as a documentation, include it or not doesn't matter
EXPOSE 8080
```

## Build and run the image

```bash
docker build -t <image name>:<tag> .

# run the image
# -p 8080:8080 means mapping the container's 8080 port to the localhost's 8080 port
docker run -p 8080:8080 <image name>:<tag>
```

## Push the image to a registry

```bash
docker tag <image name>:<tag> <username>/<image name>
docker push <username>/<image name>
```

## Bind Mounts

When the example app receives a request with path "/text", it reads the file `./text` as a string, and return this string as a response. During development, it is convenient if the app can directly read files on your host machine without build the image everytime some files get updated. Here docker's bind mount feature comes into play.

```bash
docker run --mount type=bind,src=<host path>,target=<container path> <image name>:<tag>
```
