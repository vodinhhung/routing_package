clean:
	rm -rf bin/*
	rm -rf cmd/*.exe
	rm -rf cmd/*.out

build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main