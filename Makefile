ifndef DOCKER_VERSION
export DOCKER_VERSION=0.0.6
endif

ifndef DOCKER_NAME
export DOCKER_NAME=nguyenkhanhvn/bas-relayer
endif

# abigen:
# 	abigen --abi=./src/pkg/abi/CrossChainBridge/CrossChainBridge.abi --bin=./src/pkg/abi/CrossChainBridge/CrossChainBridge.bin --type=CrossChainBridge --pkg=abi --lang=go --out=./src/pkg/abi/cross_chain_bridge.go
# 	abigen --abi=./src/pkg/abi/RelayHub/RelayHubAlpha.abi --bin=./src/pkg/abi/RelayHub/RelayHubAlpha.bin --type=RelayHubAlpha --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub_alpha.go
# 	abigen --abi=./src/pkg/abi/RelayHub/RelayHubBeta.abi --bin=./src/pkg/abi/RelayHub/RelayHubBeta.bin --type=RelayHubBeta --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub_beta.go
# 	abigen --abi=./src/pkg/abi/RelayHub/RelayHub.abi --bin=./src/pkg/abi/RelayHub/RelayHub.bin --type=RelayHub --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub.go

abigen:
	abigen --abi=./src/pkg/abi/CrossChainBridge/CrossChainBridge.abi --type=CrossChainBridge --pkg=abi --lang=go --out=./src/pkg/abi/cross_chain_bridge.go
	abigen --abi=./src/pkg/abi/RelayHub/RelayHubAlpha.abi --type=RelayHubAlpha --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub_alpha.go
	abigen --abi=./src/pkg/abi/RelayHub/RelayHubBeta.abi --type=RelayHubBeta --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub_beta.go
	abigen --abi=./src/pkg/abi/RelayHub/RelayHub.abi --type=RelayHub --pkg=abi --lang=go --out=./src/pkg/abi/relay_hub.go

build-proofwriter:
	cd src; \
	go mod tidy; \
	go build -o ../deployment/proofwriter cmd/proofwriter/main.go

build-relayer:
	cd src; \
	go mod tidy; \
	go build -o ../deployment/relayer cmd/relayer/main.go

docker-build-relayer:
	docker build -t ${DOCKER_NAME}:${DOCKER_VERSION} .

docker-push-relayer:
	docker push ${DOCKER_NAME}:${DOCKER_VERSION}