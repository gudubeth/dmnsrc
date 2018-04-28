package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/ozgio/dmnsrc/pkg/checker"
	"github.com/ozgio/dmnsrc/pkg/input"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the vailability of domains",
	Long: `|
Checks the availability of domains. It needs full domain names. If you want
extended search with domain name generation, see the "search" command. 

For multiple domain checks, seperate domain names with comma or space

Examples:
	dmn check example.com
	dmn check example.com example.org

`,
	Args: cobra.MinimumNArgs(1),
	Run:  runCheckCmd,
}

var showWhois = false

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVarP(&showWhois, "whois", "w", false, "Show whois information of domain name")
}

func runCheckCmd(cmd *cobra.Command, args []string) {
	allNames := []string{}
	for _, arg := range args {
		names := input.GrabDomainNames(arg)
		allNames = append(allNames, names...)
	}

	for _, name := range allNames {
		info, err := checker.Whois(name)

		if err != nil {
			color.Red("❗ %s: error (%s)", name, err.Error())
		} else if strings.Contains(info, "No match") ||
			strings.Contains(info, "No entries") ||
			strings.Contains(info, "NOT FOUND") { //TODO revisit whois availiblty check
			color.Green("✔ %s: available", name)
		} else {
			color.Yellow("✘ %s: unavailable", name)
		}

		if err == nil && showWhois {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println(info)
			fmt.Println(strings.Repeat("=", 80))
		}
	}
}
