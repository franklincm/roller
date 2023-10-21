//nolint:errcheck
package cmd

import (
	"os"

	"github.com/franklincm/roller/pkg/finder"
	"github.com/spf13/cobra"
)

var permissions bool

var rootCmd = &cobra.Command{
	Use:   "roller",
	Short: "Fuzzy Finder for GCP Roles and Permissions",
	Long: `roller provides a convenient interface for fuzzy finding
GCP Role and Permissions. On first run it will fetch the complete
list of Roles and Permissons and persist it locally at:

~/.config/roller/roles.json

`,

	Run: func(cmd *cobra.Command, args []string) {
		if permissions {
			finder.FindPermissions()
		} else {
			finder.FindRoles()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&permissions, "permissions", "p", false, "find permissions for a given role")
}
