package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/cgwillz/funnl/funnls"
)

func init() {

}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "download a service to your Database",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "reddit":
			o := funnls.RdAuth()
			funnls.RdGetPosts(o)
		case "web":
			if len(args) < 2 {
				panic("You must provide a browser that was specified in Funnl's config")
			}
			funnls.WbGet(args[1])
		}

	},
}
