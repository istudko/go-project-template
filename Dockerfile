FROM golang:1.13.1 AS builder
WORKDIR /go/src/gitlab.com/listenfield/lf-machine-booking/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lf-machine-booking .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /go/src/gitlab.com/listenfield/lf-machine-booking/lf-machine-booking .
EXPOSE 80
CMD ["./lf-machine-booking", "start"]