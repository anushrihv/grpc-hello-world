FROM golang:1.14 AS builder

WORKDIR /app

COPY . .

RUN go build

FROM golang:1.14

WORKDIR /app

COPY --from=builder /app/hello-world .

EXPOSE 8080

CMD /app/hello-world