# systax=docker/dockerfile:1.4
FROM golang:1.19-alpine3.16 AS builder
ARG GIT_COMMIT
ARG VERSION
WORKDIR /build
COPY . ./
RUN go mod download && \
    go build -ldflags "-X main.Version=${VERSION} -X main.GitCommit=${GIT_COMMIT}" -o /build/ -v ./...


FROM alpine:3.16.2 AS runtime
WORKDIR /
COPY --from=builder /build/gss-api /gss-api
ENTRYPOINT [ "/gss-api" ]
