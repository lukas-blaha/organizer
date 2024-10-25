FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o workoutApp ./cmd/api

RUN chmod +x /app/workoutApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/workoutApp /app

RUN apk add --no-cache tzdata

ENV TZ=Europe/Prague

CMD [ "/app/workoutApp" ]
