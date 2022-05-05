package main

// prevents go mod tidy from removing dependencies in build ignore files

import (
	_ "github.com/pojntfx/html2goapp/pkg/converter"
	_ "github.com/srwiley/oksvg"
	_ "github.com/srwiley/rasterx"
	_ "golang.org/x/net/html"
)
