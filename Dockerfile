FROM golang:1.15.0-alpine3.12

WORKDIR /go/src/app/

RUN if [ $GO_ENV=development ]; \
    then \
        apk add git; \
        go get -u github.com/cosmtrek/air; \
    fi

COPY . .
