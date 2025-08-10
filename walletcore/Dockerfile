FROM golang:1.23
WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

# Install migrate with MySQL driver
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/walletcore/main.go"]