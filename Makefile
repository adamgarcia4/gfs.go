build-master:
	@go build -o bin/master cmd/master/master.go
run-master: build-master
	@./bin/master
proto-master:
	protoc ./api/proto/master.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative	
proto: proto-master
	@echo "Generated Proto defs"
test:
	@go test -v ./...
