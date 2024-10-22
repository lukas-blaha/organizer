FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o reminderApp ./cmd/app

RUN chmod +x /app/reminderApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/reminderApp /app

RUN apk add --no-cache tzdata

ENV TZ=Europe/Prague

CMD [ "/app/reminderApp" ]
