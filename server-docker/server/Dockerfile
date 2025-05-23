# syntax=docker/dockerfile:1

# Compile
FROM golang:1.23-alpine3.21 AS compiler
WORKDIR /app
COPY server-docker/server/go.work.template go.work
COPY common common
COPY server server
RUN mkdir build
RUN cp -r server/resources build/resources && rm -rf build/resources/windows && rm -rf build/resources/unix
RUN mkdir -p build/resources/certificates
RUN go build -ldflags="-s -w" -o build/server ./server
# Compress
FROM alpine:3.21 AS compressor
RUN apk add --no-cache upx
WORKDIR /app
COPY --from=compiler /app/build/server .
RUN upx --lzma server
# Execute
FROM scratch
EXPOSE 443
VOLUME /app/resources/config
VOLUME /app/resources/certificates
VOLUME /app/logs
WORKDIR /app
COPY --from=compiler /app/build/resources resources
COPY --from=compressor /app/server .
ENTRYPOINT ["./server"]