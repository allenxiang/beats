package fileout

import (
	"fmt"

	"github.com/elastic/beats/libbeat/logp"
)

type config struct {
	Index         string `config:"index"`
	Path          string `config:"path"`
	Filename      string `config:"filename"`
	RotateEveryKb int    `config:"rotate_every_kb" validate:"min=1"`
	NumberOfFiles int    `config:"number_of_files"`
}

var (
	defaultConfig = config{
		NumberOfFiles: 7,
		RotateEveryKb: 10 * 1024,
	}
)

func (c *config) Validate() error {
	if c.Filename == "" && c.Index == "" {
		return fmt.Errorf("File logging requires filename or index being set.")
	}

	if c.NumberOfFiles < 2 || c.NumberOfFiles > logp.RotatorMaxFiles {
		return fmt.Errorf("The number_of_files to keep should be between 2 and %v",
			logp.RotatorMaxFiles)
	}

	return nil
}
