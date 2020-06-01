// +build ignore

package main

import (
	"fmt"
	"github.com/tomkennedy513/password-gen/pkg/generator"
	"github.com/tomkennedy513/password-gen/pkg/types"
)

func main() {
	config, err := types.NewPasswordGenerationConfig(
		types.SetPasswordLength(500),
		types.SetCharacterSet([]rune("pweih234")),
		)
	if err != nil {
		panic(err)
	}
	passwordGenerator := generator.NewPasswordGenerator(config)
	generate, _ := passwordGenerator.Generate()
	fmt.Println(generate)
}
