lint:
	golangci-lint run --allow-parallel-runners --no-config --deadline=10m --enable=deadcode --enable=revive --enable=varcheck --enable=structcheck --enable=gocyclo --enable=errcheck --enable=gofmt --enable=goimports --enable=misspell --enable=unparam --enable=nakedret --enable=bodyclose --enable=gosec --enable=megacheck

test:
	go test -tags=unit -v --cover ./normalizations

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/normalizer main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/normalizer main.go

run:
	./bin/normalizer < sampleData/sample.csv > output.csv