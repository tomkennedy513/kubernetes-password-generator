package types

type Option func(*PasswordGenerationConfig) error

func SetPasswordLength(length int) Option {
	return func(config *PasswordGenerationConfig) error {
		config.Length = length
		return nil
	}
}

func SetCharacterSet(characterSet CharacterSet) Option {
	return func(config *PasswordGenerationConfig) error {
		config.CharacterSet = characterSet
		return nil
	}
}