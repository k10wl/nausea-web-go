FROM golang:1.22.0-bookworm

WORKDIR /app

COPY . .
RUN go build -o ./main ./main.go

CMD "./main"
