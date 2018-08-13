FROM golang:1.10 AS builder

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/OnGridSystems/LoadBalancer
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .
RUN chmod +x /app

FROM alpine
COPY --from=builder /app .
ENTRYPOINT ["./app"]
