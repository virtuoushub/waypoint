package main

import (
	"os"

	"github.com/hashicorp/waypoint/internal/cli"
)

func main() {
	cli.ExposeDocs = true
	os.Exit(cli.Main([]string{"waypoint", "docs"}))
}
