/*
Copyright © 2024 Just-Goo
*/
package cmd

import (
	"github.com/Just-Goo/study-pal/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new studypal database and table",
	Long:  `Initialize a new studypal database and table`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
