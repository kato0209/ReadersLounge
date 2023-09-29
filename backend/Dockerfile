FROM --platform=$BUILDPLATFORM golang:1.21.1 as builder
WORKDIR /readers-lounge/backend
COPY go.mod go.sum ./
RUN go mod download
RUN --mount=type=bind,source=./,target=./ \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o /main main.go


FROM alpine:latest as go
COPY --from=builder /main /
RUN apk update && \
    apk upgrade && \
    apk add bash git postgresql-client && \
    rm -rf /var/cache/apk/*
ENTRYPOINT ["/main"]

