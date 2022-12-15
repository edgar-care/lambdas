AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:= nlp
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= lambda.py
BUILD_DIR					:= builds
SOURCE_DIR					:= ./cmd/main
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:= app.py handlers/
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)
PACKAGE_NAME                := packages
z 							?=

all: install

build:
	@pip3 install -r requirements-lambda.txt --target $(PACKAGE_NAME) --upgrade

clean: clean-local clean-lambda

clean-local:
	@find . -type f -name '*.py[co]' -delete -o -type d -name __pycache__ -delete
	@rm -rf $(BUILD_DIR)

clean-lambda:
	@rm -rf $(PACKAGE_NAME)

ifeq ($(strip $(z)),)
deploy: zip
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(ARCHIVE_NAME).zip --output yaml
else
deploy:
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(z) --output yaml
endif

install:
	@pip3 install -r requirements-local.txt

start:
	@uvicorn app:app --reload --port 5000

zip: build
	@mkdir -p $(BUILD_DIR)
	@cd $(PACKAGE_NAME) && zip -r ../$(ARCHIVE_NAME).zip . && cd ..
	@zip -r $(ARCHIVE_NAME).zip $(BINARY_NAME) $(EXTRA_FILES)
	@zip -r $(ARCHIVE_NAME) $(BINARY_NAME) $(EXTRA_FILES)


.PHONY: all \
		build \
		clean \
		clean-local \
		clean-lambda \
		deploy \
		install \
		start \
		zip
