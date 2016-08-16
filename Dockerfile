FROM golang:1.7-alpine
RUN apk add -U bash && rm -f /var/cache/apk/*
ADD . /go/src/github.com/tmc/grpc-intro
RUN go install -v github.com/tmc/grpc-intro/...
CMD grpc-intro
EXPOSE 7000

