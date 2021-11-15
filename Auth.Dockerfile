############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git && apk add make
WORKDIR /auth
COPY . /auth

RUN make tidy
RUN make mock-prepare
RUN make mock
RUN CGO_ENABLED=0 go test ./...
WORKDIR /auth/cmd/auth
RUN go build -o /go/bin/AuthServer

############################
# STEP 2 build a small image
############################
FROM alpine
RUN mkdir /server
COPY --from=builder /auth/config.yaml /server/config.yaml
COPY --from=builder /go/bin/AuthServer /server
WORKDIR /server
EXPOSE 8088
CMD ["/server/AuthServer"]
