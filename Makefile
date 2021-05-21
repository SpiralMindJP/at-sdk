# コマンド定義
GOCMD = go
PROTOCCMD = protoc

# ディレクトリー
PROTO_DIR = proto
GRPC_DIR = go/pb

# Core サービス
CORE_PROTO_DIR = $(PROTO_DIR)/core
CORE_GRPC_DIR = $(GRPC_DIR)/core
core_proto_files = $(wildcard $(CORE_PROTO_DIR)/*.proto)
core_grpc_files = $(patsubst $(CORE_PROTO_DIR)/%.proto,$(CORE_GRPC_DIR)/%.pb.go,$(core_proto_files))

.PHONY: all core-grpc FORCE

all: core-grpc

FORCE:

#
# Core サービスの gRPC コードの生成
#
$(CORE_GRPC_DIR)/%.pb.go: $(CORE_PROTO_DIR)/%.proto
	$(PROTOCCMD) \
		-I $(CORE_PROTO_DIR) \
		--go_out=$(CORE_GRPC_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(CORE_GRPC_DIR) --go-grpc_opt=paths=source_relative \
		$^

core-grpc: $(common_grpc_files) $(core_grpc_files)
