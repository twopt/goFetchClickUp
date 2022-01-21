package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twopt/clickup/utils"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of a Clickup workspace",
	Long:  `logout allows the user to delete the access token for accessing a Clickup workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleted authentication token")
		viper.Set("token", "")
		err := viper.WriteConfigAs(utils.GetConfigFile())
		if err != nil {
			log.Fatalf("cannot write config to %s: %v", utils.GetConfigFile(), err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
