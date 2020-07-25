package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "setup funnl based on your local config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init Funnl")
	},
}
