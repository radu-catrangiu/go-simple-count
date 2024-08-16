FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o ./bin ./cmd/simple-counter

EXPOSE 8080

CMD ["./bin/simple-counter"]
