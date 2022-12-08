AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:= graphql
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= main
BUILD_DIR					:= builds
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:= schema.graphql .env
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)

all: install

build:
	@go build -o $(BINARY_NAME)

deploy: zip
	@AWS_REGION=$(AWS_REGION) aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(ARCHIVE_NAME).zip

clean:
	@rm -rf $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)

install:
	@go mod tidy

start:
	@go run main.go

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
