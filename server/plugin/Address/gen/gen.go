package main

import (
	"gorm.io/gen"
	"path/filepath" //go:generate go mod tidy
	//go:generate go mod download
	//go:generate go run gen.go
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Address/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "Address", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(
		new(model.Address),
	)
	g.Execute()
}
