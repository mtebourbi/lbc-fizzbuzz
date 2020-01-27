IMAGE_TAG ?= lbc/fizzbuzz

all: docker-build

# Run all packages tests.
test: fmt vet
	go test ./... -coverprofile output/cover.out

bench: fmt vet test
	go test ./... -bench .

# Build fizzbuzz server binary (use this if you want the docker version).
build: test
	go build -o output/bin/fizzbuzz-srv github.com/mtebourbi/lbc-fizzbuzz/cmd/fizzbuzz-srv/

docker-build: test
	docker build . -t ${IMAGE_TAG}

# Run the service locally (used for dev mode).
run: fmt vet
	go run cmd/fizzbuzz-srv/main.go

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Clean the build files.
clean:
	rm -Rf output/*