FROM golang:1.13-buster as build

WORKDIR /go/src/authorizater
ADD . .

RUN go mod download
RUN go build -o /go/main

FROM gcr.io/distroless/base-debian10
WORKDIR /go/
COPY --from=build /go/main .
COPY .env .

ENV POSTGRES_URL "host=192.168.0.116 port=5432 user=postgres password=Hello2008 dbname=postgres sslmode=disable"

EXPOSE 8080

CMD ["./main"]
