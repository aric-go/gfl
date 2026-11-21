package cmd

import (
	"gfl/utils"
	"gfl/utils/strings"
	"os"

	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:     "bump [type]",
	Aliases: []string{"bp"},
	Short:   "Bump npm version and commit package.json (no push) (alias: bp)", // Will be updated after strings load
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flagType := ""
		if len(args) > 0 {
			flagType = args[0]
		}
		flagPreid, _ := cmd.Flags().GetString("preid")

		bumpType, preid, err := utils.ResolveBumpArgs(flagType, flagPreid)
		if err != nil {
			utils.Error(err.Error())
			os.Exit(1)
		}

		if err := utils.RunBump(bumpType, preid); err != nil {
			utils.Error(err.Error())
			os.Exit(1)
		}
		utils.Success(strings.GetPath("bump.success"))
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)
	bumpCmd.Flags().String("preid", "", "prerelease identifier (eg: beta/rc)")
}
