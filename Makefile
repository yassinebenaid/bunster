preview-docs:
	cd ./docs && npm run docs:dev

build-docs:
	cd ./docs && npm run docs:build

build:
	@[ -d bin ] || mkdir bin
	@go build -o ./bin/ryuko ./cmd/ryuko

compile: build
	@rm -rf .test-build && mkdir .test-build
	@./bin/ryuko build script.test.bash -o ./bin/script --build-space .test-build
