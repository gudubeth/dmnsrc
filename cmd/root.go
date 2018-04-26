package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmn",
	Short: "dmn is tool for searching domain names",
	Long:  `Domain search tool for command line. Fo further documentation go to github.com/ozgio/dmnsrc`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
