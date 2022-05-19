
APP_NAME=goapppf

VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.0.0")
COMMIT=$(shell git rev-parse --short HEAD || echo "HEAD")
MODULE=$(shell grep ^module go.mod | awk '{print $$2;}')
LD_FLAGS="-w -X $(MODULE)/server.Version=$(VERSION) -X $(MODULE)/server.Commit=$(COMMIT)"

run: binary
	@DEV=1 ./temp/$(APP_NAME)

wasm: pfarchive
	@mkdir -p server/web
	@rm -rf server/web/app.wasm
	@GOARCH=wasm GOOS=js go build -o server/web/app.wasm -ldflags $(LD_FLAGS) cmd/main.go

binary: wasm
	@mkdir -p temp
	@go build -o temp/$(APP_NAME) -ldflags $(LD_FLAGS) cmd/main.go

deploy: binary
	scp temp/$(APP_NAME) goservice:/tmp
	ssh goservice sudo /tmp/$(APP_NAME) -action deploy

npminstall:
	@go run scripts/npminstall/main.go

pfarchive: npminstall
	@go run scripts/pfarchive/main.go

pflogos: npminstall
	@go run scripts/pflogos/main.go

tsinspect: npminstall
	@go run scripts/tsinspect/main.go

tsconvert: tsinspect
	@go run scripts/tsconvert/main.go
