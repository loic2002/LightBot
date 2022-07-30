FROM golang:1.18.4-alpine3.16 AS build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o bot ./cmd/main.go

# ----------------------------------------

FROM alpine:3.16 AS final

WORKDIR /app
COPY --from=build /build/ ./bot
ENTRYPOINT ["./bot"]
CMD ["-c", "config.json"] 