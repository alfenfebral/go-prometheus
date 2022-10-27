FROM golang:1.17

ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git

# Create work directory
RUN mkdir -p $GOPATH/src/book-service

# Set the current working directory inside the container 
WORKDIR $GOPATH/src/book-service

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# Copy the source from the current directory to the working Directory inside the container 
COPY . $GOPATH/src/book-service

RUN go get ./...

RUN go build -o main $GOPATH/src/book-service/main.go

EXPOSE 5555

CMD $GOPATH/src/book-service/main