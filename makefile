GO_WORKSPACE := .

protoc:
	protoc --go_out=$(GO_WORKSPACE) --go_opt=paths=source_relative --go-grpc_out=$(GO_WORKSPACE) --go-grpc_opt=paths=source_relative proto/*.proto
	@echo "Protoc compiler finished"