FROM golang:1.17-alpine AS builder

WORKDIR /src/backend
COPY . .

ENV GO111MODULE=on

RUN go mod tidy && go mod vendor
RUN cd app && go build -o main

FROM alpine:latest

COPY --from=builder /src/backend/app /src/backend/app
COPY --from=builder /src/backend/config /src/backend/config

WORKDIR /src/backend/app

EXPOSE 8080

CMD ["./main"]