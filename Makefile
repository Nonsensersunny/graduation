PROGRAM=graduation

.PHONY prepare build run

build: 
	go build -o bin/${PROGRAM} graduation/cmd

prepare:
	export GOPROXY=https://goproxy.io/
	export GO111MODULE=auto
	go mod vendor
	cd web2; npm install

run:
	cd web2; npm run build
	go run cmd/main.go
	export GO111MODULE=on

fast:
	go run cmd/main.go

clean:
	rm -rf ${PROGRAM}

rebuild: clean build

vue-run:
	cd web2; npm run serve

vue-clean:
	rm -rf web/.nuxt/dist/
	cp web/.nuxt/dist/client/* templates

vue-build:
	cd web; npm run build