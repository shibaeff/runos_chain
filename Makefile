build-all:
	GOOS=linux GOARCH=amd64 go build -o ./build/checkersd-linux-amd64 ./cmd/runos_chaind/

do-checksum:
	cd build && sha256sum \
		checkersd-linux-amd64 checkersd-linux-arm64 \
		checkersd-darwin-amd64 checkersd-darwin-arm64 \
		> checkers_checksum

build-with-checksum: build-all do-checksum

docker-build:
	docker run --rm -it \
		-v /home/golanger/runos_chain:/runos_chain \
		-w /runos_chain \
		golang:1.18.7 \
		make build-all