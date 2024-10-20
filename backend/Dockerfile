FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o backendApp ./cmd/api

RUN chmod +x /app/backendApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/backendApp /app

EXPOSE 8080

CMD [ "/app/backendApp" ]
