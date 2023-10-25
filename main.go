package main

import (
	"fmt"
	"github.com/AlexsJones/cli/cli"
	"github.com/AlexsJones/cli/command"
	"strings"
)

func main() {
	c := cli.NewCli()

	c.AddCommand(command.Command{Name: "encrypt", Help: "Encrypt string", Func: func(args []string) {
		fmt.Println("Select hex or base64")
	}, SubCommands: []command.Command{
		{
			Name: "hex",
			Help: "Encrypt with output hex",
			Func: func(args []string) {
				if len(args) == 0 {
					fmt.Println("No input text")
				}

				msg := strings.Join(args, " ")

				encrypt(msg, EncryptionHex)
			},
		},
		{
			Name: "base64",
			Help: "Encrypt with output base64",
			Func: func(args []string) {
				if len(args) == 0 {
					fmt.Println("No input text")
				}

				msg := strings.Join(args, " ")

				encrypt(msg, EncryptionBase64)
			},
		},
	},
	})

	c.AddCommand(command.Command{Name: "decrypt", Help: "Decrypt string", Func: func(args []string) {
		fmt.Println("Select hex or base64")
	}, SubCommands: []command.Command{
		{
			Name: "hex",
			Help: "Decrypt with input hex",
			Func: func(args []string) {
				if len(args) == 0 {
					fmt.Println("No input text")
				}

				decrypt(args[0], DecryptionHex)
			},
		},
		{
			Name: "base64",
			Help: "Decrypt with input base64",
			Func: func(args []string) {
				if len(args) == 0 {
					fmt.Println("No input text")
				}

				decrypt(args[0], DecryptionBase64)
			},
		},
	},
	})

	c.Run()
}
