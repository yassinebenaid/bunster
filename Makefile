preview-docs:
	cd ./docs && npm run docs:dev

build-docs:
	cd ./docs && npm run docs:build

build:
	@[ -d bin ] || mkdir bin
	@go build -o ./bin/bunster ./cmd/bunster

compile: build
	@rm -rf .test-build && mkdir .test-build
	@./bin/bunster build script.test.bash -o ./bin/script --build-space .test-build

dump-ast: build
	@./bin/bunster ast script.test.bash
