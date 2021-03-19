FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/polygnosics-frontend

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++ git curl lsof
RUN go mod tidy

RUN cd $GOPATH/src/polygnosics-frontend/ && go build main.go
RUN chmod 0766 $GOPATH/src/polygnosics-frontend/scripts/init.sh

# This container exposes port 8085 to the outside world
EXPOSE 8085

# Run the executable
CMD ["./scripts/init.sh"]