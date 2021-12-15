install-deps:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install github.com/envoyproxy/protoc-gen-validate
	go install github.com/kazegusuri/go-proto-zap-marshaler/protoc-gen-zap-marshaler
	go install github.com/mitchellh/protoc-gen-go-json
	brew tap bufbuild/buf
	brew install buf