package configs

import (
	"io/fs"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	API API `yaml:"api"`
	DB  DB  `yaml:"db"`
}

type (
	API struct {
		Host string `yaml:"host"`
	}

	DB struct {
		ConnSql string `yaml:"conn"`
		Migrate string `yaml:"migrate"`
	}
)

func New(path string) (config Config, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return config, errors.Wrapf(err, "open config by path %s", path)
	}
	defer func(err error) {
		multierr.AppendInto(&err, file.Close())
	}(err)

	return config, errors.Wrap(
		yaml.NewDecoder(file).Decode(&config),
		"decode config information",
	)
}
