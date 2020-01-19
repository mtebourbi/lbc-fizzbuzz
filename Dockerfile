# Build the bizzfuzz binary using golang docker image.
FROM golang:1.13.6 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/ cmd/
COPY pkg/ pkg/

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o fizzbuzz-srv github.com/mtebourbi/lbc-fizzbuzz/cmd/fizzbuzz-srv/

# Create docker image for the generated binary
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/fizzbuzz-srv .
ENTRYPOINT ["/fizzbuzz-srv"]
