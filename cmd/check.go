package cmd

import (
	"fmt"
	"regexp"
	"strings"

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

Example:
	dmn check ozgur.io
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
			fmt.Printf("%s: currently only .com, .net, .edu domains are available for search\n", name)
			continue
		}
		srv := checker.SelectRandomWhoisServer()
		info, err := checker.Whois(name, srv)
		if err != nil {
			fmt.Printf("%s: error (%s)\n", name, err.Error())
		} else if strings.Contains(info, "No match") {
			fmt.Printf("%s: available\n", name)
		} else {
			fmt.Printf("%s: not available\n", name)
		}
	}
}
