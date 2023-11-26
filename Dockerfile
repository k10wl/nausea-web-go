FROM golang:1.21.4-bookworm

WORKDIR /app

COPY . .
RUN go build -o ./main ./main.go

CMD "./main"
