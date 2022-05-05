//go:build !wasm

package server

import (
	"archive/zip"
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	brotli "github.com/anargu/gin-brotli"
	"github.com/gin-gonic/gin"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//go:embed web/*
var webDirectory embed.FS

func Run() (shutdownFunc func(ctx context.Context) error, err error) {

	address := os.Getenv("ADDRESS")
	if address == "" {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}
		address = "localhost:" + port
	}

	var listener net.Listener
	if listener, err = net.Listen("tcp4", address); err != nil {
		return nil, err
	}

	if isDevelopment() {
		fmt.Printf("running on http://%s\n", address)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(gin.Logger(), gin.Recovery(), brotli.Brotli(brotli.DefaultCompression))

	goAppHandler := gin.WrapH(BuildHandler())
	goAppRoutes := []string{"/", "/app.js", "/app.css", "/manifest.webmanifest", "/wasm_exec.js", "/app-worker.js"}
	for _, route := range goAppRoutes {
		engine.GET(route, goAppHandler)
	}

	embedHandler := gin.WrapH(http.FileServer(http.FS(webDirectory)))
	engine.GET("/web/app.wasm", embedHandler)
	engine.GET("/web/logo-192.png", embedHandler)
	engine.GET("/web/logo-512.png", embedHandler)

	engine.NoRoute(servePatternFly)
	engine.RedirectTrailingSlash = false

	server := &http.Server{Handler: engine}

	go func() {
		serveErr := server.Serve(listener)
		if serveErr != nil && serveErr != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	return server.Shutdown, nil
}

func servePatternFly(c *gin.Context) {

	name := strings.TrimPrefix(c.Request.RequestURI, "/web/")
	name = strings.TrimPrefix(name, "/")

	headers := c.Writer.Header()
	etag := c.Request.Header.Get("If-None-Match")
	if etag == getRuntimeVersion() {
		c.Writer.WriteHeader(http.StatusNotModified)
		return
	}

	open, err := webDirectory.Open("web/patternfly.zip")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer open.Close()

	buff := &bytes.Buffer{}
	zipBytes, err := io.Copy(buff, open)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	reader, err := zip.NewReader(bytes.NewReader(buff.Bytes()), zipBytes)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, err := reader.Open(name)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	headers.Set("Cache-Control", "no-cache")
	headers.Set("ETag", fmt.Sprintf("%q", getRuntimeVersion()))
	c.Writer.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(name)))
	_, err = io.Copy(c.Writer, file)

	if err != nil {
		log.Println(err)
		return
	}

}

func BuildHandler() *app.Handler {
	return &app.Handler{
		Author:      "TODO",
		Description: "go-app starter",
		Name:        "go-app starter",
		Scripts:     []string{},
		Icon: app.Icon{
			AppleTouch: "/web/logo-192.png",
			Default:    "/web/logo-192.png",
			Large:      "/web/logo-512.png",
		},
		AutoUpdateInterval: autoUpdateInterval(),
		ShortName:          "starter",
		Version:            getRuntimeVersion(),
		Styles: []string{
			"/web/patternfly.css",
		},
		Title: "go-app starter",
	}
}
