// +build ignore

package main

import (
	"fmt"
	"github.com/tomkennedy513/password-gen/pkg/password"
	"github.com/tomkennedy513/password-gen/pkg/types"
)

func main() {
	config, err := types.NewPasswordGenerationConfig(
		password.SetPasswordLength(500),
		password.SetCharacterSet([]rune("pweih234")),
		)
	if err != nil {
		panic(err)
	}
	passwordGenerator := password.NewPasswordGenerator(config)
	generate, _ := passwordGenerator.Generate()
	fmt.Println(generate)

	lw := cache.NewListWatchFromClient()
}
