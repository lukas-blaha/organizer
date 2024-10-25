FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o notifierApp ./cmd/api

RUN chmod +x /app/notifierApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/notifierApp /app

EXPOSE 8080

CMD [ "/app/notifierApp" ]
