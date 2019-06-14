GOOS?=linux
GOARCH?=amd64

GCP_PROJECT=videocoin-network

NAME=mlistener
VERSION=$$(git describe --abbrev=0)-$$(git rev-parse --short HEAD)

.PHONY: deploy

default: build

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

deps:
	env GO111MODULE=on go mod vendor
	cp -r $(GOPATH)/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1 \
	vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/

docker-build:
	docker build -t gcr.io/${GCP_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	docker push gcr.io/${GCP_PROJECT}/${NAME}:${VERSION}

dbm-status:
	goose -dir migrations -table ${NAME} postgres "${DBM_MSQLURI}" status

release: docker-build docker-push

deploy:
	ENV=${ENV} deploy/deploy.sh

