FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o luciana-user .

EXPOSE 8080

CMD ["./luciana-user"]