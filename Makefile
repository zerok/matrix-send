all: bin/matrix-send

bin:
	mkdir -p bin

bin/matrix-send: $(shell find . -name '*.go') go.mod
	cd cmd/matrix-send && go build -o ../../$@

clean:
	rm -rf bin
