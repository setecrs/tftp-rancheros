FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app .
FROM alpine
COPY --from=builder /go/bin/app /app
COPY templates /templates
COPY config.json.example /config/config.json
EXPOSE 80
CMD ["/app"]
