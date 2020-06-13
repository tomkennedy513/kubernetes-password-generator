package password

import (
	"fmt"
)

type CharacterSet []rune

type GenerationConfig struct {
	Length int
	CharacterSet CharacterSet
}

func NewPasswordGenerationConfig(options ...Option) (*GenerationConfig, error) {
	config := &GenerationConfig{}
	for _, option := range options {
		err := option(config)
		if err != nil {
			return nil, fmt.Errorf("could not apply password config option -> %s", err.Error())
		}
	}
	return config, nil
}
