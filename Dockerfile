# systax=docker/dockerfile:1.4
# golang:1.19-alpine3.16
FROM golang@sha256:d475cef843a02575ebdcb1416d98cd76bab90a5ae8bc2cd15f357fc08b6a329f AS builder
ARG BINARY="gss-api"
WORKDIR /build
COPY . ./
RUN go mod download && \
    go build -v -o /build/${BINARY} ./...


# alpine:3.16.2
FROM alpine@sha256:bc41182d7ef5ffc53a40b044e725193bc10142a1243f395ee852a8d9730fc2ad AS runtime
ARG BINARY="gss-api"
COPY --from=builder /build/${BINARY} /svc
CMD [ "/svc" ]
