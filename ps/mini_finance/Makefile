# Basic Config

PROTO_SRC = proto
PROTO_FILES = $(shell find $(PROTO_SRC) -name '*.proto')

# Main Task

.PHONY: all generate

all: generate

generate:
	@echo "🚀 Generating all .proto files..."
	@for file in $(PROTO_FILES); do \
		echo "➡️  Compiling $$file"; \
		protoc \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			$$file; \
	done
	@echo "✅ Selesai generate semua .proto!"

