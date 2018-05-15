FROM golang:latest AS builder
ADD server.go /
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM busybox
COPY --from=builder /server /
CMD ["/server"]