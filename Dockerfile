FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go build -o /app/out ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/out /root/
CMD ["/root/out"]
