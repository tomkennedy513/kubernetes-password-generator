package generator

import (
	"crypto/rand"
	"fmt"
	"github.com/tomkennedy513/password-gen/pkg/types"
	"math/big"
)

const DefaultPasswordLength int = 32

var DefaultCharacterSet types.CharacterSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Generator interface {
	Generate() (string, error)
}

type generator struct {
	config *types.PasswordGenerationConfig
}

func NewPasswordGenerator(config *types.PasswordGenerationConfig) Generator {
	return &generator{
		config: config,
	}
}

func (g *generator) Generate() (string, error) {
	var (
		length       int
		characterSet types.CharacterSet
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
