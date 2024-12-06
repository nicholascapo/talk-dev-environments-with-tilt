FROM golang:1.23-alpine as tools
RUN go install github.com/air-verse/air@latest

################################
FROM golang:1.23-alpine as development

WORKDIR /go/src/example

CMD ["air", "-c", "/go/src/example/air.toml"]
COPY --from=tools /go/bin/air /go/bin/air
COPY . .

################################
FROM golang:1.23-alpine AS builder

WORKDIR /go/src/
COPY . .
RUN go build -o /go/bin/example .

################################
FROM gcr.io/distroless/static-debian11:nonroot as production

ENTRYPOINT ["/go/bin/example"]
COPY --from=builder /go/bin/example /go/bin/example
