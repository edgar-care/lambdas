AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:=
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= main
BUILD_DIR					:= builds
SOURCE_DIR					:= ./cmd/main
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:=
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)
GO_BINARY_PATH  			:= $(shell go env GOPATH)/bin

all: install

build:
	@go build -o $(BINARY_NAME) $(SOURCE_DIR)

deploy: zip
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(ARCHIVE_NAME).zip --output yaml

clean:
	@rm -rf $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)

install:
	@go mod tidy

start:
	@$(GO_BINARY_PATH)/air -c .air.toml

zip: build
	@mkdir -p $(BUILD_DIR)
	@zip builds/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE) $(BINARY_NAME) $(EXTRA_FILES)


.PHONY: all \
		build \
		clean \
		deploy \
		install \
		start \
		zip
