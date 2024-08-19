FROM golang:1.21-alpine AS builder

WORKDIR /src/laundry
COPY . .

ENV GO111MODULE=on

RUN go mod tidy && go mod vendor
RUN cd app && go build -o main

FROM alpine:latest

RUN apk update && apk add curl

COPY --from=builder /src/laundry/app /src/laundry/app
COPY --from=builder /src/laundry/config /src/laundry/config
COPY --from=builder /src/laundry/migrations /src/laundry/migrations

WORKDIR /src/laundry/app