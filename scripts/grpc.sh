protoc --proto_path=social-media-proto \
    --go_out=proto \
    --go_opt=paths=source_relative \
    --go-grpc_out=proto \
    --go-grpc_opt=paths=source_relative \
    social-media-proto/base.proto \
    social-media-proto/like.proto \
    social-media-proto/config.proto