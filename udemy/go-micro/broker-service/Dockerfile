# the container where the built code will be executed should be smallest as possible
FROM alpine:latest
RUN mkdir /app
COPY broker /app
CMD ["app/broker"]

# From here on, it shows the previous code that built the app using a builder image then ran it on another image
# Now, the app is being built by the Makefile commands

# FROM golang:1.18 AS builder

# RUN mkdir /app

# # copy all local content to container folder "app"
# COPY . /app

# WORKDIR /app

# # compiles /cmd/api files and output as "broker" executable
# RUN CGO_ENABLED=0 go build -o broker ./cmd/api

# RUN chmod +x /app/broker

# # the container where the built code will be executed should be smallest as possible
# FROM alpine:latest
# RUN mkdir /app
# COPY --from=builder /app/broker /app
# CMD ["app/broker"]