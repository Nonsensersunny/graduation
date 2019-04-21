PROGRAM=graduation
DOCKER_TARGET=hub.deepin.io/deepin/mirrors_status
DOCKER_BUILD_TARGET=${DOCKER_TARGET}.builder

build: 
	go build -o ${PROGRAM} graduation/cmd

prepare:
	export GOPROXY=https://goproxy.io/
	export GO111MODULE=auto

run: prepare
	go run cmd/main.go

docker:
	docker build -f deployments/Dockerfile --target builder -t ${DOCKER_BUILD_TARGET} .
	docker build -f deployments/Dockerfile -t ${DOCKER_TARGET} .

docker-push:
	docker push ${DOCKER_BUILD_TARGET}
	docker push ${DOCKER_TARGET}

clean:
	rm -rf ${PROGRAM}

rebuild: clean build

vue-run:
	cd web
	npm run dev

vue-build:
	cd web
	npm run build
	npm start