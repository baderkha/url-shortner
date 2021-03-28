start:
	make build-local
	cd dist && ./main

build-local:prepare-build-dir
	cp env.local.json dist/env.json
	go build -o dist/main src/main.go
build-deploy:prepare-build-dir
	cp env.json dist/env.json
	GOOS=linux go build -o dist/main src/main.go

deploy:build-deploy
	serverless deploy
	
prepare-build-dir:
	mkdir -p dist/client
	cp README.md dist/README.md
	cp -R client/* dist/client/

build:prepare-build-dir
	cp env.json dist/env.json

	
