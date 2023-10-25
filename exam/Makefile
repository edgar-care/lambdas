AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:= exam
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= main
BUILD_DIR					:= builds
SOURCE_DIR					:= ./cmd/main
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:= .env
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)
GO_BINARY_PATH  			:= $(shell go env GOPATH)/bin
GOOS						:= linux
GOARCH						:= amd64
z 							?=

all: install

build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(BINARY_NAME) $(SOURCE_DIR)

ifeq ($(strip $(z)),)
deploy: zip
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(ARCHIVE_NAME).zip --output yaml
else
deploy:
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(z) --output yaml
endif

clean:
	@rm -rf $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@rm -rf tmp/

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
