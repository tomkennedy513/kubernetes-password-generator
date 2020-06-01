FROM golang:alpine as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o password-gen .

FROM alpine:latest as runner
WORKDIR /app
COPY --from=builder /app/password-gen .
ENTRYPOINT ["./password-gen"]