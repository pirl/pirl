# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: pirl android ios pirl-cross swarm evm all test clean
.PHONY: pirl-linux pirl-linux-386 pirl-linux-amd64 pirl-linux-mips64 pirl-linux-mips64le
.PHONY: pirl-linux-arm pirl-linux-arm-5 pirl-linux-arm-6 pirl-linux-arm-7 pirl-linux-arm64
.PHONY: pirl-darwin pirl-darwin-386 pirl-darwin-amd64
.PHONY: pirl-windows pirl-windows-386 pirl-windows-amd64
##export GOPATH=$(pwd)
GOBIN = $(shell pwd)/build/bin
GO ?= latest

pirl:
	build/env.sh go run build/ci.go install ./cmd/pirl
	@echo "Done building."
	@echo "Run \"$(GOBIN)/pirl\" to launch pirl."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/pirl.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/pirl.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

pirl-cross: pirl-linux pirl-darwin pirl-windows pirl-android pirl-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/pirl-*

pirl-linux: pirl-linux-386 pirl-linux-amd64 pirl-linux-arm pirl-linux-mips64 pirl-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-*

pirl-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/pirl
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep 386

pirl-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/pirl
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep amd64

pirl-linux-arm: pirl-linux-arm-5 pirl-linux-arm-6 pirl-linux-arm-7 pirl-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep arm

pirl-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/pirl
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep arm-5

pirl-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/pirl
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep arm-6

pirl-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/pirl
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep arm-7

pirl-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/pirl
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep arm64

pirl-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/pirl
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep mips

pirl-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/pirl
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep mipsle

pirl-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/pirl
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep mips64

pirl-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/pirl
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/pirl-linux-* | grep mips64le

pirl-darwin: pirl-darwin-386 pirl-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/pirl-darwin-*

pirl-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/pirl
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-darwin-* | grep 386

pirl-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/pirl
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-darwin-* | grep amd64

pirl-windows: pirl-windows-386 pirl-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/pirl-windows-*

pirl-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/pirl
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-windows-* | grep 386

pirl-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/pirl
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/pirl-windows-* | grep amd64
