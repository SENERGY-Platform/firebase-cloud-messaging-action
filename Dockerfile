FROM golang:1.21 AS builder

COPY . /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM scratch
COPY --from=builder /go/src/app/app .
CMD ["./app"]
