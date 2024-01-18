FROM app-build-base:0.1.0 AS builder
WORKDIR /go/src/
COPY ./ /go/src/
RUN go build -o ../bin/res ./cmd/

FROM ubuntu:22.04 AS runner
WORKDIR /app
COPY --from=builder /go/bin/res ./
ENTRYPOINT ["./res"]
