user_proto:
	rm -rf ./user_info/pb/*.go
	protoc --proto_path=./user_info/proto --go_out=./user_info/pb --go_opt=paths=source_relative \
    --go-grpc_out=./user_info/pb --go-grpc_opt=paths=source_relative \
    ./user_info/proto/*.proto


login_proto:
	rm -rf ./login/pb/*.go
	protoc --proto_path=./login/proto --go_out=./login/pb --go_opt=paths=source_relative \
    --go-grpc_out=./login/pb --go-grpc_opt=paths=source_relative \
    ./login/proto/*.proto