
APP_NAME=goapppf

VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.0.0")
COMMIT=$(shell git rev-parse --short HEAD || echo "HEAD")
MODULE=$(shell grep ^module go.mod | awk '{print $$2;}')
LD_FLAGS="-w -X $(MODULE)/server.Version=$(VERSION) -X $(MODULE)/server.Commit=$(COMMIT)"

run: binary
	@DEV=1 ./temp/$(APP_NAME)

wasm: pf_archive
	@mkdir -p server/web
	@rm -rf server/web/app.wasm
	@GOARCH=wasm GOOS=js go build -o server/web/app.wasm -ldflags $(LD_FLAGS) cmd/main.go

binary: wasm
	@mkdir -p temp
	@go build -o temp/$(APP_NAME)	-ldflags $(LD_FLAGS) cmd/main.go

deploy: binary
	scp temp/$(APP_NAME) goservice:/tmp
	ssh goservice sudo /tmp/$(APP_NAME) -action deploy

pf_archive:
	go run scripts/genpfarchive.go
	cp temp/node_pf/patternfly.zip server/web
	go run scripts/genpflogos.go

htmltogo:
	go run scripts/htmltogo.go




