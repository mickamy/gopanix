package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

func String() string {
	if version == "dev" {
		return "dev (built from source; try GitHub release for accurate version)"
	}
	return version
}

var Cmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show gopanix version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gopanix version %s\n", String())
	},
}
