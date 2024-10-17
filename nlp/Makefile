AWS_REGION 					:= eu-west-3
AWS_LAMBDA_FUNCTION_NAME	:= nlp
AWS_CLI_PROFILE				:= edgar.care
BINARY_NAME					:= lambda.py
BUILD_DIR					:= builds
SOURCE_DIR					:= .
DATE						:= $(shell date +"%d%m%y_%H%M%S")
EXTRA_FILES					:= app.py handlers/ .env
ARCHIVE_NAME				:= $(BUILD_DIR)/$(AWS_LAMBDA_FUNCTION_NAME)$(DATE)
PACKAGE_NAME                := 146778342232.dkr.ecr.eu-west-3.amazonaws.com
REQUIREMENTS_FILE			:= requirements-lambda.txt
ZIP_FILE					:= test.zip
z 							?=
# VM_USER 					= $(VM_USER_ENV)
# VM_HOST 					= $(VM_HOST_ENV)
# VM_PORT 					= $(VM_PORT_ENV)
# VM_DIR 						= /home/$(VM_USER)/nlp
# ARCHIVE_NAME 				= deploy_package.tar.gz
# DOCKER_IMAGE 				= edgar/nlp-local
# DOCKER_CONTAINER_NAME 		= nlp_app

all: install

build:
	@(docker build -f Dockerfile.lambda -t $(AWS_LAMBDA_FUNCTION_NAME):latest .)

tag:
	@docker tag $(AWS_LAMBDA_FUNCTION_NAME):latest $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest

clean: clean-local 

clean-local:
	@find . -type f -name '*.py[co]' -delete -o -type d -name __pycache__ -delete
	@rm -rf $(BUILD_DIR)

ifeq ($(strip $(z)),)
deploy: zip
	@aws ecr get-login-password --region eu-west-3 | docker login --username AWS --password-stdin $(PACKAGE_NAME)
	@docker build -f Dockerfile.lambda -t $(AWS_LAMBDA_FUNCTION_NAME):latest .
	@docker tag $(AWS_LAMBDA_FUNCTION_NAME):latest $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest
	@docker push $(PACKAGE_NAME)/$(AWS_LAMBDA_FUNCTION_NAME):latest
else
deploy:
	@AWS_REGION=$(AWS_REGION) AWS_PAGER= aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION_NAME) --profile $(AWS_CLI_PROFILE) --zip-file fileb://$(z) --output yaml
endif

# deploy:
# 	tar --exclude=$(ARCHIVE_NAME) -czvf $(ARCHIVE_NAME) .

# 	scp -P $(VM_PORT) $(ARCHIVE_NAME) $(VM_USER)@$(VM_HOST):$(VM_DIR)

# 	ssh -p $(VM_PORT) $(VM_USER)@$(VM_HOST) "\
# 		cd $(VM_DIR) && \
# 		tar -xzvf $(ARCHIVE_NAME) && \
# 		rm $(ARCHIVE_NAME) && \
# 		sudo docker build -t $(DOCKER_IMAGE) . && \
# 		sudo docker stop $(DOCKER_CONTAINER_NAME) || true && \
# 		sudo docker rm $(DOCKER_CONTAINER_NAME) || true && \
# 		sudo docker run -d -p 5999:5999 --name $(DOCKER_CONTAINER_NAME) -v $(VM_DIR):/code $(DOCKER_IMAGE) \
# 	"

install:
	@pip3 install -r requirements-local.txt
	@python3 -m spacy download fr_core_news_lg

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
