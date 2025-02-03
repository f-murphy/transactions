FROM golang:alpine

RUN apk add --no-cache postgresql-client

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o bank ./cmd/main.go

RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]