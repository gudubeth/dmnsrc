package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates domain names",
	Long: `Generates domain names without checking for their availability. 
This is the same mechanism used by "search" command internally. Example:

dmn generate 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
