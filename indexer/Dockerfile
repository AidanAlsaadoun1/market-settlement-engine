# ---- Build stage -----------------------------------------------------
# Multi-stage build: the toolchain never ships. The final image carries
# a single static binary on a minimal base, which keeps the image small
# and the CVE surface near zero.
FROM golang:1.22-alpine AS build

WORKDIR /src

# Dependencies first, so Docker's layer cache survives code changes:
# go.mod/go.sum rarely change, source changes constantly.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO off -> fully static binary, safe on a bare minimal base.
RUN CGO_ENABLED=0 go build -o /indexer .

# ---- Runtime stage ---------------------------------------------------
FROM alpine:3.20

# TLS roots for HTTPS calls to the RPC endpoint; nothing else needed.
RUN apk add --no-cache ca-certificates

# Never run as root, even in a toy.
RUN adduser -D -u 10001 indexer
USER indexer

COPY --from=build /indexer /usr/local/bin/indexer

ENTRYPOINT ["indexer"]