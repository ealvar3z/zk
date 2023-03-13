package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "zk",
	Short: "personal zettlekasten system",
}

func NewRootCommand() *cobra.Command { return rootCmd }

func init() {
	rootCmd := NewRootCommand()
	rootCmd.AddCommand(add)
}
