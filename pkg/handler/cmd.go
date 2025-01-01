package handler

import (
	"fmt"

	ghclient "github.com/milinddethe15/profilefetch/internal/github"
	"github.com/milinddethe15/profilefetch/pkg/display"
	"github.com/spf13/cobra"
)

func RunCmd(cmd *cobra.Command, args []string) {
	username := args[0]
	client := ghclient.CreateClient()
	user, err := client.FetchUserProfile(username)
	if err != nil {
		fmt.Printf("Error fetching profile for %s: %v\n", username, err)
		return
	}
	display.DisplayOutput(username, user)
}
