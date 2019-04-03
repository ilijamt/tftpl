package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tftpl",
	Short: "Terraform File Templater",
	Long:  `Templates terraform files based on yaml data`,
}

func Execute() error  {
	return rootCmd.Execute()
}