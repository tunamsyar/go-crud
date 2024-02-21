FROM golang:1.21-alpine

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY *.go ./

RUN go build -o /go-webserver

EXPOSE 8080

CMD ["/go-webserver"]

