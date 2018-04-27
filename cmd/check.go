package cmd

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/ozgio/dmnsrc/pkg/checker"
	"github.com/ozgio/dmnsrc/pkg/input"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check availability of domains",
	Long: `|
Checks the availability of domains. It needs full domain names and doesn't 
generate any domains. For extended search see "search" command. 

For multiple domain check seperate domains with comma or space

Examples:
	dmn check example.com
	dmn check example.com example.org

`,
	Args: cobra.MinimumNArgs(1),
	Run:  runCheckCmd,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func runCheckCmd(cmd *cobra.Command, args []string) {
	allNames := []string{}
	for _, arg := range args {
		names := input.GrabDomainNames(arg)
		allNames = append(allNames, names...)
	}

	var validSuffix = regexp.MustCompile(`\.com|\.net|\.edu$`)
	for _, name := range allNames {
		if !validSuffix.Match([]byte(name)) {
			color.Red("❗ %s: error (currently only .com, .net, .edu domains are available for search)", name)
			continue
		}
		srv := checker.SelectRandomWhoisServer()
		info, err := checker.Whois(name, srv)
		if err != nil {
			color.Red("❗ %s: error (%s)", name, err.Error())
		} else if strings.Contains(info, "No match") {
			color.Green("✔ %s: available", name)
		} else {
			color.Yellow("✘ %s: not available", name)
		}
	}
}
