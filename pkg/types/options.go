package types

import "github.com/tomkennedy513/password-gen/pkg/generator"

type Option func(config *generator.PasswordGenerationConfig) error

func SetPasswordLength(length int) Option {
	return func(config *generator.PasswordGenerationConfig) error {
		config.Length = length
		return nil
	}
}

func SetCharacterSet(characterSet generator.CharacterSet) Option {
	return func(config *generator.PasswordGenerationConfig) error {
		config.CharacterSet = characterSet
		return nil
	}
}