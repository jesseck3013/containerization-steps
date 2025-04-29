FROM golang

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

# serve as a documentation, include it or not doesn't matter
EXPOSE 8080


type=bind,src=.,target=/usr/src/app
