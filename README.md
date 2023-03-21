# Bas-relayer
Bas-relayer contain 2 part: ***proofwriter*** and ***relayer***


**Proofwriter**

Use to generate proofs manually.

**Relayer**

Automatic check and transfer proofs from source chain to destination chain.

## **Deploy**
**Docker**

Install **Docker** and run:
```
DOCKER_NAME=myrelayer DOCKER_VERSION=0.0.1 make docker-build-relayer
```

**Manual**

*1. Installing all the dependencies of project*

**Bas-relayer** write in golang. To install all the dependencies of project, install ```golang``` and, run:
```sh
go mod tidy
```

*2.1. Build **Proofwriter***
```
make build-proofwriter
```

*2.2. Build **Relayer***
```
make build-relayer
```

*3. Config*

Create ```./deployment/config.json``` file based on ```./deployment/config-template.json```

Or use make file:
```sh
cd deployment
make create-relayer-config
```

*4. Run **Relayer***

```sh
cd deployment
make start-relayer
```