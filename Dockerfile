FROM golang:1.22.2-bullseye

WORKDIR /app

COPY go.mod .
RUN go mod download && go mod verify

COPY . .
RUN go build -o /app/main .
CMD ["/app/main"]