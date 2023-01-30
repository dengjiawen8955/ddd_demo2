make_dir:=$(shell pwd)
app_name:=$(shell basename $(make_dir))

# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen tidy build run

## exec.sql: Create dadabase and import sql
.PHONY: exec_sql
exec.sql:
	cd tools/exec_sql/ && go run . && cd $(make_dir)

## exec.sql.force: Drop database and create dadabase and import sql
.PHONY: exec.sql.force
exec.sql.force:
	cd tools/exec_sql/ && go run . -f true && cd $(make_dir)

## gen: Gemerate protobuf files.
.PHONY: gen
gen:
	cd internal/servers/rpc/proto_file/ && \
	protoc -I ./ --go_out=./ --go-grpc_out=./ ./in/* && \
	rm -rf ../protos && cp -r protos/ ../ && rm -rf protos && \
	cd $(make_dir)
	# 安装 gojson 工具
	cd internal/common/tools/gojson/ && ./install.sh
	cd $(make_dir)
	# 生成 json 文件
	gojson internal/ "*_dto.go" json/

## tidy: Tidy go mod.	
.PHONY: tidy
tidy:
	go mod tidy

## build: Build app
.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o ./bin/$(app_name) ./main.go

## run: Run app
.PHONY: run
run: build
	./bin/$(app_name) --config ./config.yaml

## up: Docker compose up
.PHONY: up
up: build
	docker-compose  -f docker-compose.yaml up --build -d

## up.env: Docker compose up environment
.PHONY: up.env
up.env:
	docker-compose  -f docker-compose-env.yaml up --build -d

## help: Show this help info.
.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
