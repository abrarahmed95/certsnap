package certsnap

import (
	"sync"

	"github.com/abrarahmed95/certsnap/internal/certsnap"
	"github.com/spf13/cobra"
)

var (
	wg      sync.WaitGroup
	jsonOut bool
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Check expiry of the specified domains",
	Run: func(cmd *cobra.Command, args []string) {
		certsnap.ValidateAndCheckCertificates(args, &wg, jsonOut)
	},
}

func init() {
	checkCmd.Flags().BoolVarP(&jsonOut, "json", "j", false, "Print results in JSON format")
	rootCmd.AddCommand(checkCmd)

}
