package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var tokenLength int

var generateTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate a secure random token",
	Long:  `Generates a cryptographically secure random token, useful for secrets, API keys, or session IDs.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := generateToken(tokenLength)
		if err != nil {
			fmt.Println("Error generating token:", err)
			return
		}
		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(generateTokenCmd)
	generateTokenCmd.Flags().IntVarP(&tokenLength, "length", "l", 32, "Length of token in bytes (will be doubled in hex)")
}

func generateToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
