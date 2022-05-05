//go:build wasm

package server

import (
	"context"
)

func Run() (shutdownFunc func(ctx context.Context) error, err error) { return }
