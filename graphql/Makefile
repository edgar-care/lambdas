all: install

build:
	@yarn build

clean:
	@rm -rf node_modules/
	@rm -rf build/

install:
	@yarn install

deploy:
	@npx prisma db push
	@echo "No deploy specified"

start:
	@yarn start

prisma:
	@npx prisma generate

.PHONY: all \
		build \
		install \
		deploy \
		start \
		prisma
