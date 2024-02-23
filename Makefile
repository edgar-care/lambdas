DIRECTORIES	:= $(shell ls -d */)
t ?=
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)


all: install


ifeq ($(strip $(t)),)
build:
	@for d in $(DIRECTORIES); do \
		make -C $${d} build; \
	done

install:
	@for d in $(DIRECTORIES); do \
		make -C $${d} install; \
	done

.PHONY: build \
		deploy \
		install
else
build:
	@make -C $(t) build

deploy:
	@make -C $(t) deploy

deploy-stage:
	@make -C $(t) deploy-stage stage=$(stage)

install:
	@make -C $(t) install

start:
	@make -C $(t) start

test:
	@make -C $(t) test

.PHONY: build \
		deploy \
		install \
		start	\
		test
endif


ifeq ($(BRANCH), dev)
pr:
	@gh pr create --base main --fill
else
pr:
	@gh pr create --base dev --fill
	@echo "\033[0;31;1mMake sure that your lambda is deployed !!\033[0m"
endif


.PHONY: all \`
		pr
