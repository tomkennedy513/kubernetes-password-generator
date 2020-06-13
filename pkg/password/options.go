package password

type Option func(config *GenerationConfig) error

func SetPasswordLength(length int) Option {
	return func(config *GenerationConfig) error {
		if length != 0 {
			config.Length = length
		}
		return nil
	}
}

func SetCharacterSet(characterSet CharacterSet) Option {
	return func(config *GenerationConfig) error {
		if len(characterSet) != 0 {
			config.CharacterSet = characterSet
		}
		return nil
	}
}