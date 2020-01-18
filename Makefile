

test:
	go test ./... -v

build:
	go build -o output/bin/fizzbuzz-srv github.com/mtebourbi/lbc-fizzbuzz/cmd/fizzbuzz-srv/

clean:
	rm -Rf output