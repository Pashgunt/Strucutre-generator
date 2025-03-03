package services

import (
	"github.com/Pashgunt/Strucutre-generator/internal/config"
	"slices"
	"strings"
)

type ConfigParser struct {
	file      string
	structure string
	mode      string
}

func (c *ConfigParser) File() string {
	return c.file
}

func (c *ConfigParser) SetFile(file string) *ConfigParser {
	c.file = file

	return c
}

func (c *ConfigParser) Structure() string {
	return c.structure
}

func (c *ConfigParser) SetStructure(structure string) *ConfigParser {
	c.structure = structure

	return c
}

func (c *ConfigParser) Mode() string {
	return c.mode
}

func (c *ConfigParser) SetMode(mode string) *ConfigParser {
	lowerMode := strings.ToLower(mode)

	if !slices.Contains([]string{config.CliMode, config.FileMode}, lowerMode) {
		c.mode = config.CliMode
	} else {
		c.mode = lowerMode
	}

	return c
}
