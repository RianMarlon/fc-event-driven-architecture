FROM golang:1.23
WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/walletcore/main.go"]