# syntax=docker/dockerfile:1

# Compile
FROM golang:1.23-alpine3.21 AS compiler
WORKDIR /app
COPY server-docker/genCert/go.work.template go.work
COPY common common
COPY server-genCert server-genCert
RUN mkdir -p build/resources/certificates
RUN mkdir build/bin
RUN go build -ldflags="-s -w" -o build/bin/genCert ./server-genCert
# Compress
FROM alpine:3.21 AS compressor
RUN apk add --no-cache upx
WORKDIR /app
COPY --from=compiler /app/build .
RUN upx --lzma bin/genCert
# Execute
FROM scratch
VOLUME /app/resources/certificates
WORKDIR /app
COPY --from=compiler /app/build/resources resources
COPY --from=compressor /app .
ENTRYPOINT ["./bin/genCert"]