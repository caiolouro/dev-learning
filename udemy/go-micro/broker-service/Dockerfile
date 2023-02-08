FROM golang:1.18 AS builder

RUN mkdir /app

# copy all local content to container folder "app"
COPY . /app

WORKDIR /app

# compiles /cmd/api files and output as "broker" executable
RUN CGO_ENABLED=0 go build -o broker ./cmd/api

RUN chmod +x /app/broker

# the container where the built code will be executed should be smallest as possible
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/broker /app
CMD ["app/broker"]