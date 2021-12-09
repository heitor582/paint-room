FROM golang:1.17-alpine

WORKDIR /go/src/back-paint-room

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN GOOS=linux go build -o /back-paint-room

EXPOSE 8080

CMD ["/back-paint-room"]