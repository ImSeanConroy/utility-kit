package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var generatePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a secure random password",
	Long:  `Generates a secure password using uppercase letters, lowercase letters, and numbers.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword()
		fmt.Println(password)
	},
}

var capitals = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lowers = []rune("abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("0123456789")

func generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	patterns := [][]([]rune){
		{capitals, capitals, lowers},
		{numbers, lowers, numbers},
		{lowers, capitals, capitals},
		{numbers, lowers, lowers},
	}

	var parts []string

	for _, pattern := range patterns {
		var section []rune
		for _, pool := range pattern {
			section = append(section, pool[rand.Intn(len(pool))])
		}
		parts = append(parts, string(section))
	}

	return strings.Join(parts, "-")
}
