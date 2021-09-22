FROM golang:1.17
WORKDIR /go/delivery
COPY . .
RUN go mod download
CMD go run cmd/server/main.go