FROM golang:1.7-alpine
RUN apk add -U bash && rm -f /var/cache/apk/*
ADD . /go/src/github.com/tmc/go-helloworld
RUN go install github.com/tmc/go-helloworld
CMD go-helloworld
EXPOSE 7000

