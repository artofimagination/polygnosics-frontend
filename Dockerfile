FROM golang:1.15.2-alpine

WORKDIR $GOPATH/src/polygnosics-frontend
ARG SERVER_PORT

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN apk add --update g++ git curl lsof
RUN go mod tidy

RUN cd $GOPATH/src/polygnosics-frontend/ && go build main.go
RUN chmod 0766 $GOPATH/src/polygnosics-frontend/scripts/init.sh

# This container exposes SERVER_PORT to the outside world.
# Check .env for the actual value
EXPOSE $SERVER_PORT

# Run the executable
CMD ["./scripts/init.sh"]