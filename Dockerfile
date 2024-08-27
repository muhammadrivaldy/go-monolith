FROM golang:1.21-alpine AS builder

WORKDIR /src/backend
COPY . .

ENV GO111MODULE=on

RUN go mod tidy && go mod vendor
RUN cd app && go build -o main

FROM alpine:latest

RUN apk update && apk add curl

COPY --from=builder /src/backend/app /src/backend/app
COPY --from=builder /src/backend/config /src/backend/config
COPY --from=builder /src/backend/migrations /src/backend/migrations

WORKDIR /src/backend/app