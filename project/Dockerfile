FROM golang:1.15.8-alpine3.13

WORKDIR /app

COPY . /app

RUN go build server/main.go

EXPOSE 3000

ENTRYPOINT ["./main"]
