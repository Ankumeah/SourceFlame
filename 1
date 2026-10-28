FROM golang:1.27.3-alpine3.24@sha256:b8c9b82affbbb82a020bbcf998dad2ba9516944ac7e883524642a481a9505e9c AS builder

WORKDIR /opt/build/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build \
    -o build/out \
    -ldflags='-s -w' \
    ./cmd/sourceflame/

FROM alpine@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

WORKDIR /opt/sf/

RUN apk add git --no-cache

COPY --from=builder /opt/build/build/out sf

ENTRYPOINT [ "/opt/sf/sf" ]
