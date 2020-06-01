package generator

import (
	"fmt"
	"github.com/tomkennedy513/password-gen/pkg/types"
)

type CharacterSet []rune

type PasswordGenerationConfig struct {
	Length int
	CharacterSet CharacterSet
}

func NewPasswordGenerationConfig(options ...types.Option) (*PasswordGenerationConfig, error) {
	config := &PasswordGenerationConfig{}
	for _, option := range options {
		err := option(config)
		if err != nil {
			return nil, fmt.Errorf("could not apply password config option -> %s", err.Error())
		}
	}
	return config, nil
}
