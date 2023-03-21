ifndef TESTNET_TRANSACTION_HASH
export TESTNET_TRANSACTION_HASH=""
endif
ifndef BAS_TRANSACTION_HASH
export BAS_TRANSACTION_HASH=""
endif
ifndef MAINNET_TRANSACTION_HASH
export MAINNET_TRANSACTION_HASH=""
endif
ifndef LOCALHOST_TRANSACTION_HASH
export LOCALHOST_TRANSACTION_HASH=""
endif

ifndef TESTNET_BLOCK_NUMBER
export TESTNET_BLOCK_NUMBER=0
endif
ifndef BAS_BLOCK_NUMBER
export BAS_BLOCK_NUMBER=1300000
endif
ifndef MAINNET_BLOCK_NUMBER
export MAINNET_BLOCK_NUMBER=0
endif
ifndef LOCALHOST_BLOCK_NUMBER
export LOCALHOST_BLOCK_NUMBER=0
endif

ifndef INTERVAL
export INTERVAL=0
endif

ifndef DATA_PATH
export DATA_PATH="./data/"
endif

.PHONY: gen-block-proofs-testnet
gen-block-proofs-testnet:
	envsubst < config-template.json > config.json; \
	./proofwriter --network=testnet genblockproofs

.PHONY: gen-block-proofs-bas
gen-block-proofs-bas:
	envsubst < config-template.json > config.json; \
	./proofwriter --network=bas genblockproofs

.PHONY: gen-block-header-testnet
gen-block-header-testnet:
	envsubst < config-template.json > config.json; \
	./proofwriter --network=testnet genheader

.PHONY: gen-block-header-bas
gen-block-header-bas:
	envsubst < config-template.json > config.json; \
	./proofwriter --network=bas genheader

.PHONY: create-relayer-config
create-relayer-config:
	envsubst < config-template.json > config.json
	cp config.json ../docker-compose

.PHONY: start-relayer
start-relayer:
	envsubst < config-template.json > config.json; \
	./relayer
