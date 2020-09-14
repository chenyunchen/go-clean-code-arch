install-mock:
	brew install vektra/tap/mockery
	brew upgrade mockery

gen-mock:
	cd internal/pkg/domain; mockery --all

BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -mod=vendor -o ${BINARY} *.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t sample .

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint