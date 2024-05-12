FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /wd
COPY . /wd
RUN go get -d -v
RUN go build

FROM scratch
ARG version=0.1.0
ARG release=unreleased
LABEL name="player-Go" \
      maintainer="Fachschaft IT, HS-Esslingen" \
      version=${version} \
      release=${release} \
      summary="A client for BitWars written in Go using gin-gonic" \
      description="A client for BitWars written in Rust using the gin-gonic webserver"
EXPOSE 3000
COPY --from=builder /wd/player-Go /
CMD ["./player-Go"]
