FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/dummy-backend
ARG SERVER_PORT

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++ git curl lsof
RUN go mod tidy

RUN cd $GOPATH/src/dummy-backend/ && go build main.go

# This container exposes port 8082 to the outside world
EXPOSE $SERVER_PORT

# Run the executable
CMD ["./main"]