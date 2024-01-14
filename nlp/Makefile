AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:= nlp
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= lambda.py
BUILD_DIR					:= builds
SOURCE_DIR					:= ./cmd/main
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:= app.py handlers/ .env
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)
PACKAGE_NAME                := 146778342232.dkr.ecr.eu-west-3.amazonaws.com
z 							?=

all: install

build:
	@(docker build -f Dockerfile.lambda -t $(AWS_LAMBDA_FUNCTION_NAME):latest .)

tag:
	@docker tag $(AWS_LAMBDA_FUNCTION_NAME):latest $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest

clean: clean-local clean-lambda

clean-local:
	@find . -type f -name '*.py[co]' -delete -o -type d -name __pycache__ -delete
	@rm -rf $(BUILD_DIR)

clean-lambda:
	@rm -rf $(PACKAGE_NAME)

ifeq ($(strip $(z)),)
deploy: zip
	@aws ecr get-login-password --region eu-west-3 | docker login --username AWS --password-stdin $(PACKAGE_NAME)
	@(docker build -f Dockerfile.lambda -t $(AWS_LAMBDA_FUNCTION_NAME):latest .)
	@docker tag $(AWS_LAMBDA_FUNCTION_NAME):latest $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest
	@docker push $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest
else
deploy:
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(z) --output yaml
endif

install:
	@pip3 install -r requirements-local.txt
	@python3 -m spacy download fr_core_news_lg

# start:
# 	@uvicorn app:app --reload --port 5000

start:
	@(docker build -f Dockerfile.local -t edgar/nlp-local .) && (docker run -p 5999:5999 -v $(shell pwd):/code -it edgar/nlp-local)

zip: build
	@mkdir -p $(BUILD_DIR)
	@zip $(ARCHIVE_NAME) . $(EXTRA_FILES)

test:
	@python3 -m coverage run -m unittest discover -s tests -v
	@python3 -m coverage report
	@python3 -m coverage html

.PHONY: all \
		build \
		clean \
		clean-local \
		clean-lambda \
		deploy \
		install \
		start \
		zip \
		test
