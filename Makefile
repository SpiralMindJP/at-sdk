# コマンド定義
GOCMD = go
PROTOCCMD = protoc

# ディレクトリー
PROTO_DIR = proto
GRPC_DIR = go/pb

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

.PHONY: all common-protobuf core-protobuf FORCE

all: core-protobuf

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
