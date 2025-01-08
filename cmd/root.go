package cmd

import (
	"log"

	"github.com/milinddethe15/profilefetch/pkg/handler"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "profilefetch [username]",
	Short: "ProfileFetch - A CLI to fetch GitHub profile information",
	Long:  "ProfileFetch is a command-line tool to fetch and display GitHub profile information in a customizable and aesthetic format.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handler.RunCmd(cmd, args)
	},
}

func init() {
	RootCmd.Flags().BoolP("show-picture", "p", false, "Display profile picture in ASCII format")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
