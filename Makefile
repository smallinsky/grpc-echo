SVC_NAME ?= echo
REV = $$(git rev-parse --short HEAD)
SERVER = ${SVC_NAME}-server:${REV}
CLIENT = ${SVC_NAME}-client:${REV}
DOCKER_USER ?= msmallinsky


build-docker: build-docker-server build-docker-client

build-docker-server:
	docker build -t ${SERVER} . -f Dockerfile.server

build-docker-client:
	docker build -t ${CLIENT} . -f Dockerfile.client

publish: publish-server publish-client

publish-server:
	docker tag  ${SERVER} ${DOCKER_USER}/${SERVER}
	docker push ${DOCKER_USER}/${SERVER}

publish-client:
	docker tag  ${CLIENT} ${DOCKER_USER}/${CLIENT}
	docker push ${DOCKER_USER}/${CLIENT}
	
gen_proto:
	protoc --go_out=plugins=grpc:. ./proto/service.proto

