//go:generate packer-sdc mapstructure-to-hcl2 -type Config
package web

import (
	"os"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	PathName  	string `mapstructure:"path_name" required:"false"`
}

func (c *Config) Prepare(raws ...interface{}) ([]string, error) {
	err := config.Decode(c, &config.DecodeOpts{
		Interpolate:       true,
		InterpolateFilter: &interpolate.RenderFilter{},
	}, raws...)

	if err != nil {
		return nil, err
	}

	// If the pathname is not set, default to the home directory
	if c.PathName == "" {
		HOME := os.Getenv("HOME")
		DIR := HOME + "/web"
		err := os.Mkdir(DIR, 0755)

		if err != nil {
			return nil, err
		}
		c.PathName = DIR
	}

	return nil, nil
}