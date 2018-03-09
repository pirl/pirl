.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/DaCHRIS/Iceberg-/cmd/swarm
	go build -o bin/geth     github.com/DaCHRIS/Iceberg-/cmd/geth
	go build -o bin/bootnode github.com/DaCHRIS/Iceberg-/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/DaCHRIS/Iceberg-/swarm/...
