FROM golang:alpine AS builder

RUN apk --no-cache add ca-certificates tzdata

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY ./ ./

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .
RUN cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime

FROM scratch

COPY --from=builder /dist/main .
COPY --from=builder /dist/.env .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/localtime /etc/localtime

ENV TZ=Asia/Seoul

ENTRYPOINT ["./main"]