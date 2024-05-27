build:
	@go build -o bin/fs
run: build
	@./bin/fs
proto-master:
	protoc ./api/proto/master.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative	
proto: proto-master
	@echo "Generated Proto defs"
test:
	@go test -v ./...
