package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const DefaultPasswordLength int = 32

var DefaultCharacterSet CharacterSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()?")

type Generator interface {
	Generate() (string, error)
}

type generator struct {
	config *GenerationConfig
}

func NewPasswordGenerator(config *GenerationConfig) Generator {
	return &generator{
		config: config,
	}
}

func (g *generator) Generate() (string, error) {
	var (
		length       int
		characterSet CharacterSet
		result string
	)

	if length = g.config.Length; length == 0 {
		length = DefaultPasswordLength
	}

	if characterSet = g.config.CharacterSet; len(characterSet) == 0 {
		characterSet = DefaultCharacterSet
	}

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
		if err != nil {
			return "", fmt.Errorf("could not generate password -> could not generate random number -> %s", err.Error())
		}

		result += string(characterSet[n.Int64()])
	}

	return result, nil
}
