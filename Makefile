# Currently we are not using all deps but those tools will be introduced as we move on

.PHONY: all get-deps clean build build-all release certs test docker proto docker-clean sit service-integration-test

all: docker

# 'get-deps' fetches all go build deps
get-deps:
	@go get -u github.com/tools/godep
	@go get -u golang.org/x/tools/cmd/cover
	@go get -u github.com/golang/mock/mockgen
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/micro/protobuf/protoc-gen-go
	@go get -u github.com/jstemmer/go-junit-report
	@go get -u github.com/axw/gocov/gocov
	@go get -u github.com/AlekSi/gocov-xml
	@git submodule init
	@git submodule update

# 'clean' cleans up any binaries and docker containers
clean: docker-clean-build
	@rm -rf out

# 'docker-clean-build' cleans up any containers and images left from the build process for this service
# excluding the released builds
docker-clean-build:
	./scripts/ci/docker_clean.sh MidasLab

# 'docker-clean' removes ALL docker images and containers for this service
docker-clean: docker-clean-build
	./scripts/ci/docker_clean.sh MidasLab full

# 'static' creates a statically linked executable. It does not recompile all packages.
static:
	@export GO_SERVICE_STATIC_BUILD=true; ./scripts/ci/build.sh MidasLab

# 'build' creates a dynamically linked executable. Useful for local testing
build:
	./scripts/ci/build.sh MidasLab

# 'build-all' recompiles all packages
build-all:
	./scripts/ci/build.sh MidasLab -a

# 'test' runs all the tests
test: 
	./scripts/ci/test.sh


gen-proto:
	go generate ./tools.go

gen-mock:
	mockery --all --dir=./proto/src/go

up:
	docker-compose up -d

down:
	docker-compose down -v

build-mocks:
