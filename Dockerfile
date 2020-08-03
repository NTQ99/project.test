FROM golang:latest

WORKDIR /apps/project.test

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV HOST 0.0.0.0

RUN go build ./cmd/c-r-m-server/main.go

CMD ["./main"]