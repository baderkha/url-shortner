start:
	docker-compose -f docker/docker-compose.yml up -d
	make build-local
	cd dist && ./main
shutodwn:
	docker-compose -f docker/docker-compose.yml down
build-local:prepare-build-dir
	cp env.local.json dist/env.json
	go build -o dist/main src/main.go
prepare-build-dir:
	mkdir -p dist

build:prepare-build-dir
	cp env.json dist/env.json

	
