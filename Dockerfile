FROM golang:latest as builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux go build -o ./api ./cmd/api/
CMD ["./api"]