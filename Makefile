DIRECTORIES	:= $(shell ls -d */)
t ?=


all: install


ifeq ($(strip $(t)),)
build:
	@for d in $(DIRECTORIES); do \
		make -C $${d} build; \
	done

deploy:
	@for d in $(DIRECTORIES); do \
		make -C $${d} install; \
	done
	@terraform apply

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

install:
	@make -C $(t) install

start:
	@make -C $(t) start

.PHONY: build \
		deploy \
		install \
		start
endif


terraform:
	@terraform init


.PHONY: all \
		terraform
