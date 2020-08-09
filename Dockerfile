FROM golang:1.14

WORKDIR /app

COPY . .

RUN go build

RUN ls

RUN pwd

EXPOSE 8080

CMD /app/hello-world