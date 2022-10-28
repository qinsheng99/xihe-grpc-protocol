go ?= go
protoc ?= protoc

define tips
	$(info )
	$(info *************** $(1) ***************)
	$(info )
endef

.PHONY: protocol

protocol:
	$(call tips,Gen Protocol)
	$(protoc) -I ./proto --go_out=./protocol training.proto
	$(protoc) -I ./proto --go-grpc_out=./protocol training.proto

.PHONY: clean

clean:
	rm ./protocol
