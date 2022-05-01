FROM golang:1.17-alpine AS builder

WORKDIR /src/umkm/backend
COPY . .

ENV GO111MODULE=on

RUN go mod tidy && go mod vendor
RUN cd app && go build -o main

FROM alpine:latest

COPY --from=builder /src/umkm/backend/app /src/umkm/backend/app
COPY --from=builder /src/umkm/backend/config /src/umkm/backend/config

WORKDIR /src/umkm/backend/app

EXPOSE 8080

CMD ["./main"]