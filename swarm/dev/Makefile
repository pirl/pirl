.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/pirl/pirl/cmd/swarm
	go build -o bin/geth     github.com/pirl/pirl/cmd/geth
	go build -o bin/bootnode github.com/pirl/pirl/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/pirl/pirl/swarm/...
