.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    git.pirl.io/community/pirl/cmd/swarm
	go build -o bin/pirl     git.pirl.io/community/pirl/cmd/pirl
	go build -o bin/bootnode git.pirl.io/community/pirl/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v git.pirl.io/community/pirl/swarm/...
