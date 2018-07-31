FROM golang:1.7.3 AS builder

RUN wget -O dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \
    chmod 755 dep && \
    mv dep /usr/bin/

WORKDIR /go/src/github.com/example-project-go/
COPY . ./
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/example-project-go/server ./server-go/
EXPOSE 8005
CMD ["./server-go/server"]  