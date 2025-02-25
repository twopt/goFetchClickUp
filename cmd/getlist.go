package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/twopt/clickup/client"
	"github.com/twopt/clickup/internal"
)

var listCmd = &cobra.Command{
	Use:   "list LISTID",
	Short: "get data for a list object by supplying it's list id",
	Long: `Request JSON data for a list objectin an authorized 
	Clickup workspace`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckTokenExists(); !authed {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		l := client.ListRequest{
			ListID: strings.Trim(args[0], " "),
		}

		client.Request(l)
	},
}

func init() {
	getCmd.AddCommand(listCmd)
}
