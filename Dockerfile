FROM golang:1.20-bullseye

ADD . /app
WORKDIR /app

RUN go build -o main cmd/api/main.go
RUN chmod +x main

CMD ["./main"]