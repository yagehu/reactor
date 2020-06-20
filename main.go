package main

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor/internal/app"
)

func main() {
	fx.New(app.New()).Run()
}
