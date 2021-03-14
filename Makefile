ENV = development

build/client:
	cd ./client; yarn run build

build/server:
	cd ./server; go mod tidy;
	cd ./server/src; GOOS=linux GOARCH=amd64 go build -o ../bin/server

build:
	$(MAKE) build/client
	$(MAKE) build/server	

docker/remove:
	docker image rm workout_ap_${ENV} workout_db_${ENV}

docker/build:
	sudo docker-compose -f ./docker/docker-compose.${ENV}.yml build

docker/rebuild:
	$(MAKE) docker/remove
	$(MAKE) docker/build

docker/up:
	docker-compose -f ./docker/docker-compose.${ENV}.yml up -d

docker/down:
	docker-compose -f ./docker/docker-compose.${ENV}.yml down