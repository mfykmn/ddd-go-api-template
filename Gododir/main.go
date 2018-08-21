package main

import (
	"os"

	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	goPath := os.Getenv("GOPATH")

	p.Task("server", nil, func(c *do.Context) {
		// rebuilds and restarts when a watched file changes
		c.Start("main.go -c {{.config}}", do.M{
			"config": goPath + "/src/github.com/mafuyuk/ddd-go-api-template/_tools/local/api.toml",
			"$in":    "cmd/api/",
		})
	}).Src("*.go", "**/*.go", "**/**/*.go").Debounce(3000)
}

func main() {
	do.Godo(tasks)
}
