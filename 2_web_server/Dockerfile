FROM golang:latest

WORKDIR /usr/src/app

COPY main.go .

RUN go build -o web_server main.go

EXPOSE 8080

CMD ["./web_server"]
