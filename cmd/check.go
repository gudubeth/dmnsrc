package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ozgio/dmnsrc/pkg/dev"
	"github.com/ozgio/dmnsrc/pkg/input"
	"github.com/ozgio/dmnsrc/pkg/whois"
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
var benchmark = false

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVarP(&showWhois, "whois", "w", false, "Show whois information of domain name")
	checkCmd.Flags().BoolVarP(&benchmark, "benchmark", "b", false, "Benchmark")
}

func runCheckCmd(cmd *cobra.Command, args []string) {
	if benchmark {
		defer dev.PrintElapsedTime("Running time", time.Now())
	}
	allNames := []string{}
	for _, arg := range args {
		names := input.GrabDomainNames(arg)
		allNames = append(allNames, names...)
	}
	out := whois.FetchMultiple(allNames, 4)

	for record := range out {
		benchmarkStr := ""
		if benchmark {
			benchmarkStr = fmt.Sprintf(" (%.3f)", float32(record.Elapsed/time.Millisecond)/1000)
		}

		if record.Error != nil {
			color.Red("❗ %s: error (%s)%s", record.Name, record.Error.Error(), benchmarkStr)
		} else if record.Available {
			color.Green("✔ %s: available%s", record.Name, benchmarkStr)
		} else {
			color.Yellow("✘ %s: unavailable%s", record.Name, benchmarkStr)
		}

		if record.Error == nil && showWhois {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println(record.Response)
			fmt.Println(strings.Repeat("=", 80))
		}
	}

}
