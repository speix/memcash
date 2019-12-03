BINARY_NAME=memcash
CONTAINER_NAME=supergramm/$(BINARY_NAME):1.0

# Generate protobuf code based on on pb/*.proto files
proto:
	@echo "Building protobuf files"
	rm -rf app/pb
	protoc -I=. -I=$(GOPATH)/src --go_out=plugins=grpc:. \
	pb/$(BINARY_NAME)_messages.proto \
    pb/$(BINARY_NAME).proto

# Compile and build project to a single binary
compile:
	@echo "Building go executable"
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o docker/$(BINARY_NAME)

# Compile and build container
container:
	@echo "Building dockerfile"
	cd ./docker && docker build -t $(CONTAINER_NAME) .

# Push container in the DockerHub registry
deploy:
	@echo "Pushing to DockerHub"
	docker push $(CONTAINER_NAME)

# Run the unit tests of the project
test:
	@echo "Testing project"
	go test

# Remove object files from package source directories
clean:
	go clean
	rm -rf app/pb
	rm docker/$(BINARY_NAME)

# Run all processes at once
all: clean proto test compile container deploy