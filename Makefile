PROGRAM=graduation
TAG=1.0

all: prepare build
	@echo ${PROGRAM}


docker:
	docker build .

prepare:
	export GOPATH=${PWD}
	export PATH=${GOPATH}/bin;${PATH}
	go mod vendor
	cd web2; npm install

go-build:
	go build -o bin/${PROGRAM} ${PROGRAM}/cmd

build:
	rm -rf dist && mkdir dist
	go build -o bin/${PROGRAM} ${PROGRAM}/cmd
	cd web2; npm run build
	mkdir -p dist/html
	cp bin/* dist
	cp -r web2/dist/* dist/html
	cp -r config dist

fast:
	go run cmd/main.go

login:
	docker login

push: login build
	docker build . -t ${PROGRAM}:${TAG}
	docker push ${PROGRAM}:${TAG}
