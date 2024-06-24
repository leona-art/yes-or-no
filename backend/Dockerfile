FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go 

FROM alpine:latest

WORKDIR /app

ENV GEMINI_API_KEY="[キーを追加]"

COPY --from=builder /app/main .

CMD ["./main"]