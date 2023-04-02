FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/

COPY . .

# get dependencies.
RUN go get -d -v

RUN go build -o /surfdb

EXPOSE 8080

CMD [ "/surfdb" ]