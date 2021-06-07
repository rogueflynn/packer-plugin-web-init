package main

import (
	"fmt"
	"os"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/rogueflynn/packer-plugin-web-init/builder/web"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(web.Builder))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
