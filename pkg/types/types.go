package types

import "fmt"

type CharacterSet []rune

type PasswordGenerationConfig struct {
	Length int
	CharacterSet CharacterSet
}

func NewPasswordGenerationConfig(options ...Option) (*PasswordGenerationConfig, error) {
	config := &PasswordGenerationConfig{}
	for _, option := range options {
		err := option(config)
		if err != nil {
			return nil, fmt.Errorf("could not apply password config option -> %s", err.Error())
		}
	}
	return config, nil
}



