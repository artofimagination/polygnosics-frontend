FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/github.com/artofimagination/polygnosics-frontend

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++ git curl lsof
RUN go mod tidy

RUN cd $GOPATH/src/github.com/artofimagination/polygnosics-frontend/ && go build main.go
RUN chmod 0766 $GOPATH/src/github.com/artofimagination/polygnosics-frontend/scripts/init.sh

# This container exposes port 8081 to the outside world
EXPOSE 8081

# Run the executable
CMD ["./main"]