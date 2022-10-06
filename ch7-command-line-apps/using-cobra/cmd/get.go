package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run:   RunGet,
}

func init() {
	rootCmd.AddCommand(getCmd)

	// define flags

}

func RunGet(cmd *cobra.Command, args []string) {
	const devAddr = "127.0.0.1:3450"

	fs := cmd.Flags()
}
