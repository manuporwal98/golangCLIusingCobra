package stringer

import (
	"fmt"

	"github.com/manuporwal98/Stringer_Cobra/pkg/stringer"
	"github.com/spf13/cobra"
)

var Onlydigits bool
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := stringer.Inspect(i, Onlydigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&Onlydigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}
