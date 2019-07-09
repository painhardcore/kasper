FROM golang:1.12-alpine as builder
RUN apk add --update --no-cache openssl-dev curl git openssh gcc musl-dev linux-headers util-linux && \
  rm -rf /tmp/* /var/cache/apk/*
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o counter ./cmd/...

FROM scratch
COPY --from=builder /build/counter /app/
WORKDIR /app
CMD ["./counter"]