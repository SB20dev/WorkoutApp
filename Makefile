ENV = development

build/client:
	cd ./client; yarn run build

build/server:
	which dep || go get -v -u github.com/golang/dep/cmd/dep
	cd ./server/src; dep ensure; GOOS=linux GOARCH=amd64 go build -o ../bin/server

build:
	$(MAKE) build/client
	$(MAKE) build/server	

run:
	./server/bin/server

docker/up:
	docker-compose -f ./docker/docker-compose.${ENV}.yml up -d

docker/down:
	docker-compose -f ./docker/docker-compose.${ENV}.yml down

docker/rebuild:
	docker image rm workout_ap_${ENV} workout_db_${ENV}
	sudo docker-compose -f ./docker/docker-compose.${ENV}.yml build