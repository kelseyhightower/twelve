FROM golang:1.14rc1-buster as go-builder
ENV GO111MODULE=on
WORKDIR /module
COPY . /module/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo \
      -ldflags '-w -extldflags "-static"' \
      -o twelve

FROM scratch
COPY --from=go-builder /module/twelve .
ENTRYPOINT ["/twelve"]
