FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build -o hello_world main.go

CMD ["./hello_world"]
