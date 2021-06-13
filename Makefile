BINARY=tass-cli
APP_VERSION=v0.1.0
SCHEDULE_VERSION=v0.1.0
BUILD=`date +%FT%T%z`
GIT_COMMIT=`git rev-parse HEAD`

# The following flags are params in go build
FILE_PATH=github.com/tass-io/cli/pkg/version
APP_VERSION_FLAG=${FILE_PATH}.AppVersion=${APP_VERSION}
SCHEDULE_VERSION_FLAG=${FILE_PATH}.LocalSchedulerVersion=${SCHEDULE_VERSION}
BUILD_FLAG=${FILE_PATH}.BuildTime=${BUILD}
GIT_COMMIT_FLAG=${FILE_PATH}.GitCommit=${GIT_COMMIT}
FLAGS="-X ${APP_VERSION_FLAG} -X ${SCHEDULE_VERSION_FLAG} -X ${BUILD_FLAG} -X ${GIT_COMMIT_FLAG}"

SCHEDULER_PATH=/root/tass-io/scheduler

build: fmt vet
	@echo "  >  Building binary..."
	@go build -ldflags ${FLAGS} -o bin/${BINARY} main.go

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

docker:
	docker build -t tass-io/cli:latest .

clean:
	@echo "  >  Cleaning build cache"
	@rm bin/${BINARY}

clean-mock-fn: build
	@bin/${BINARY} function delete --name function1
	@bin/${BINARY} function delete --name function2
	@bin/${BINARY} function delete --name function3

mock-fn: clean-mock-fn
	@echo " >  Replacing the old mock zip files..."
	@cp ${SCHEDULER_PATH}/user-code/plugin-golang-wrapper.zip mock/plugin-golang-wrapper.zip
	@echo " >  Replacing the old mock zip files done."
	@echo " >  Mocking function1, function2 and function3 ..."
	@bin/${BINARY} function create --name function1 -m true
	@bin/${BINARY} function create --name function2 -m true
	@bin/${BINARY} function create --name function3 -m true

.PHONY: build fmt vet test docker clean help


