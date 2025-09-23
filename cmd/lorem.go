package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var wordCount int

var generateLoremCmd = &cobra.Command{
	Use:   "lorem",
	Short: "Generate lorem ipsum text",
	Long:  `Generates placeholder lorem ipsum text with a given number of words.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generateLorem(wordCount))
	},
}

func init() {
	rootCmd.AddCommand(generateLoremCmd)
	generateLoremCmd.Flags().IntVarP(&wordCount, "words", "w", 50, "Number of words to generate")
}

var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet",
	"consectetur", "adipiscing", "elit", "sed", "do",
	"eiusmod", "tempor", "incididunt", "ut", "labore",
	"et", "dolore", "magna", "aliqua", "ut",
	"enim", "ad", "minim", "veniam", "quis",
	"nostrud", "exercitation", "ullamco", "laboris", "nisi",
	"ut", "aliquip", "ex", "ea", "commodo",
	"consequat", "duis", "aute", "irure", "dolor",
	"in", "reprehenderit", "in", "voluptate", "velit",
	"esse", "cillum", "dolore", "eu", "fugiat",
	"nulla", "pariatur", "excepteur", "sint", "occaecat",
	"cupidatat", "non", "proident", "sunt", "in",
	"culpa", "qui", "officia", "deserunt", "mollit",
	"anim", "id", "est", "laborum",
}

func generateLorem(n int) string {
	rand.Seed(time.Now().UnixNano())
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = loremWords[rand.Intn(len(loremWords))]
	}
	words[0] = strings.Title(words[0])
	return strings.Join(words, " ") + "."
}
