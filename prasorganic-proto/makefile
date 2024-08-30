GO_MODULE := github.com/dwprz/prasorganic-proto

.PHONY: clean
clean: 
	rm --force --recursive ./protogen 
	mkdir -p ./protogen

.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/user/*.proto ./proto/user/type/*.proto \
	./proto/otp/*.proto ./proto/otp/type/*.proto \
	./proto/product/*.proto ./proto/product/type/*.proto \
	./proto/order/*.proto ./proto/order/type/*.proto \

.PHONY: build
build: clean protoc-go

.PHONY: licenses
licenses:
	rm -rf ./LICENSES
	go-licenses save ./... --save_path=./LICENSES