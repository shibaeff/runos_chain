build-all:
	GOOS=linux GOARCH=amd64 go build -o ./build/runos_chaind-linux-amd64 ./cmd/runos_chaind/

do-checksum:
	cd build && sha256sum \
		runos_chaind-linux-amd64 runos_chaind-linux-arm64 \
		runos_chaind-darwin-amd64 runos_chaind-darwin-arm64 \
		> checkers_checksum

build-with-checksum: build-all do-checksum

build-all-rest:
	GOOS=linux GOARCH=amd64 go build -o ./build/runos_chaind-rest-linux-amd64 ./cmd/runos_chaind-rest

docker-build-rest:
	docker build -f Dockerfile-runos_chaind-rest . -t runos_chaind-rest:latest --build-arg BUILDARCH=amd64

clearn-docker-data:
	for i in $(ls docker); do rm -rf docker/$i/data; done
