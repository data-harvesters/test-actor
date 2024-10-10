FROM golang:1.22.1

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o ./test-actor
CMD ["./test-actor"]