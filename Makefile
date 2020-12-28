all:
	$(MAKE) build
	$(MAKE) run
build:
	cd ./client; yarn run build
	which dep || go get -v -u github.com/golang/dep/cmd/dep
	cd ./server/src; dep ensure; GOOS=linux GOARCH=amd64 go build -o ../bin/server
run:
	./server/bin/server
