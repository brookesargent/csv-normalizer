lint:
	golangci-lint run --allow-parallel-runners --no-config --deadline=10m --enable=deadcode --enable=revive --enable=varcheck --enable=structcheck --enable=gocyclo --enable=errcheck --enable=gofmt --enable=goimports --enable=misspell --enable=unparam --enable=nakedret --enable=prealloc --enable=bodyclose --enable=gosec --enable=megacheck

test:
	go test -tags=unit -v --cover ./normalizations

build:
	go build -o bin/normalizer main.go