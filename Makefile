# コマンド定義
GOCMD = go
PROTOCCMD = protoc

# ディレクトリー
PROTO_DIR = proto
GO_DIR = go
GRPC_DIR = $(GO_DIR)/pb

# Common protobuf
COMMON_PROTO_DIR = $(PROTO_DIR)/common
COMMON_GRPC_DIR = $(GRPC_DIR)/common
common_proto_files = $(wildcard $(COMMON_PROTO_DIR)/*.proto)
common_go_files = $(patsubst $(COMMON_PROTO_DIR)/%.proto,$(COMMON_GRPC_DIR)/%.pb.go,$(common_proto_files))

# Core サービス
CORE_PROTO_DIR = $(PROTO_DIR)/core
CORE_GRPC_DIR = $(GRPC_DIR)/core
core_proto_files = $(wildcard $(CORE_PROTO_DIR)/*.proto)
core_grpc_files = $(patsubst $(CORE_PROTO_DIR)/%.proto,$(CORE_GRPC_DIR)/%.pb.go,$(core_proto_files))

# Error details
ERRDETAILS_PROTO_DIR = $(PROTO_DIR)/errdetails
ERRDETAILS_GO_DIR = $(GO_DIR)/errors/errdetails
errdetails_proto_files = $(wildcard $(ERRDETAILS_PROTO_DIR)/*.proto)
errdetails_go_files = $(patsubst $(ERRDETAILS_PROTO_DIR)/%.proto,$(ERRDETAILS_GO_DIR)/%.pb.go,$(errdetails_proto_files))

.PHONY: all common-protobuf core-protobuf errdetails-protobuf FORCE

all: core-protobuf errdetails-protobuf

FORCE:

#
# Common proto から Go コードの生成
#
$(COMMON_GRPC_DIR)/%.pb.go: $(COMMON_PROTO_DIR)/%.proto
	$(PROTOCCMD) \
		-I $(PROTO_DIR) \
		--go_out=$(GRPC_DIR) --go_opt=paths=source_relative \
		$<

common-protobuf: $(common_go_files)

#
# Core サービスの gRPC コードの生成
#
$(CORE_GRPC_DIR)/%.pb.go $(CORE_GRPC_DIR)/%_grpc.pb.go: $(CORE_PROTO_DIR)/%.proto
	$(PROTOCCMD) \
		-I $(PROTO_DIR) \
		--go_out=$(GRPC_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GRPC_DIR) --go-grpc_opt=paths=source_relative \
		$<

core-protobuf: common-protobuf $(core_grpc_files)

#
# error details proto から Go コードの生成
#
$(ERRDETAILS_GO_DIR)/%.pb.go: $(ERRDETAILS_PROTO_DIR)/%.proto
	$(PROTOCCMD) \
		-I $(PROTO_DIR) \
		--go_out=./ --go_opt=paths=import \
		$<

errdetails-protobuf: $(errdetails_go_files)
