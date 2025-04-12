FROM golang:1.24.1
RUN apk add --no-cache git curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server

EXPOSE 8080

CMD ["/server"]