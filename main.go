package main

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"

	"github.com/mach-composer/mach-composer-plugin-contentful/internal"
)

func main() {
	p := internal.NewContentfulPlugin()
	plugin.ServePlugin(p)
}
