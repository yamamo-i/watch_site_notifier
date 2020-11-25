FROM golang:1.15.5-alpine3.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/app
COPY . .
RUN go build

FROM alpine:3.12.1
RUN apk add ca-certificates
COPY --from=builder /go/src/app/watch_site_notifier /watch_site_notifier
CMD ["/watch_site_notifier"]
