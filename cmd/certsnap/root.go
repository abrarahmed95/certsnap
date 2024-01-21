/*
Copyright © 2024 Abrar Ahmed abrarahmed377@hotmail.com
*/
package certsnap

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "certsnap",
	Short: "Check SSL certificate expiration status for domains",
	Long:  `CertSnap is a cli tool to check the expiration status of SSL certificates for specified domains.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a root command")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
