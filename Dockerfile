FROM golang:1.17.2

RUN apt-get update \
&& apt-get install -y mariadb-client

WORKDIR /go/src/go_simple_api

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -o server main.go

RUN chmod +x entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]

EXPOSE 8080

CMD ["./server"]
