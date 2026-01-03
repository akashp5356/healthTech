FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN  GOOS=linux go build -o healthtech-backend ./server/main.go
EXPOSE 8080
CMD ["./healthtech-backend"]
