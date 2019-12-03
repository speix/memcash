# Memcash

Time persistent key/value memory cache gRPC service.

## Description

This is a memory storage system similar to memcached. The user is able to set a key/value pair of data into memory for a given time. The communication is built around protocol buffers and gRPC.

The app supports authentication using environment variables and exposes gRPC endpoints to add and retrieve data. It also supports health checking functionality in case it integrates with a container-orchestration system like Kubernetes.

## Installation

Before using the app you need to setup the environment variables:


```bash
MEMCASH_GRPC_PORT="The port of the gRPC server"
MEMCASH_GRPC_USER="The user for authenticated requests"
MEMCASH_GRPC_PASS="The password for authenticated requests"
MEMCASH_TICKER_INTERVAL="The clock for purging the memory"
```
Make sure that no other application is using that port.

## Usage

Select the way to run the app that suits you the most:

- Using Docker Compose tool:

```docker
docker-compose up -d
```

- Using go
```golang
go get github.com/speix/memcash
go install github.com/speix/memcash
```


## License
[MIT](https://choosealicense.com/licenses/mit/)