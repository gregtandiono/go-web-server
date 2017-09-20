FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go install -v

ENTRYPOINT /go/bin/app

EXPOSE 8080
