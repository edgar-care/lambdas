<<<<<<< HEAD
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
z 							?=

all: install

build:
	@go build -o $(BINARY_NAME) $(SOURCE_DIR)

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
=======
all: install


build:
	@echo "No need to build"


deploy:
	@terraform apply

>>>>>>> d0232d5 (feat(workflow): Makefile architecture with example lambda folder)

install:
	@go mod tidy

<<<<<<< HEAD
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
=======

start:
	@echo "I am an example lambda"


.PHONY:  build \
		install \
		start
>>>>>>> d0232d5 (feat(workflow): Makefile architecture with example lambda folder)
