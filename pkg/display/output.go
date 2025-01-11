package display

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v68/github"
	"github.com/milinddethe15/profilefetch/pkg/utility"
)

func DisplayOutput(username string, user *github.User, repos []*github.Repository, showAvatar bool) {
	// color styles
	title := color.New(color.FgHiCyan).SprintFunc()
	info := color.New(color.FgWhite).SprintFunc()
	label := color.New(color.FgHiMagenta).SprintFunc()

	asciiArt := `         .:*%@@@@@@@@@@@@%*-.         
      .:#@@@@@@@@@@@@@@@@@@@@#-.      
     =@@@@@@@@@@@@@@@@@@@@@@@@@@=     
   :#@@@@%@@@@@@@@@@@@@@@@@@%@@@@#:   
  -@@@@@* ..:*@%#*++*#%@#:.. *@@@@@-  
.-@@@@@@=                    =@@@@@@=.
.%@@@@@@#                    *@@@@@@%:
=@@@@@@+.                    .+@@@@@@+
%@@@@@%:                      :%@@@@@%
@@@@@@%.                      .%@@@@@@
@@@@@@@:                      :%@@@@@@
*@@@@@@+.                    .+@@@@@@#
:@@@@@@@+.                  .+@@@@@@@-
.*@@@%%@@@*-.            .-*@@@@@@@@*.
 .#@@@-.+@@@@@@:      :@@@@@@@@@@@@%. 
  .*@@@+.:*@@%-        -@@@@@@@@@@*.  
   .-@@@*... ..        :@@@@@@@@@-.   
     .-%@@@@@@:        :@@@@@@%-.     
        .*@@@@:        :@@@@*.        `

	if showAvatar {
		if art := avatarArt(user.GetAvatarURL()); art != "" {
			asciiArt = art
		}
	}
	topRepos := 3 // top 3 repositories
	var repoData strings.Builder
	for _, repo := range repos[:min(topRepos, len(repos))] {
		if repo.Name != nil {
			repoData.WriteString(fmt.Sprintf("%s, ", *repo.Name))
		}
	}
	RepoNames := utility.CleanString(repoData.String(), ", ")

	data := []string{
		title(username + "@github"),
		strings.Repeat("-", 30),
		fmt.Sprintf("%s%s", label("Name: "), info(user.GetName())),
		fmt.Sprintf("%s%s", label("Email: "), info(user.GetEmail())),
		fmt.Sprintf("%s%s", label("Bio: "), info(user.GetBio())),
		fmt.Sprintf("%s%s", label("Followers: "), info(fmt.Sprintf("%d", user.GetFollowers()))),
		fmt.Sprintf("%s%s", label("Following: "), info(fmt.Sprintf("%d", user.GetFollowing()))),
		fmt.Sprintf("%s%s", label("Top Repositories: "), info(RepoNames)),
	}

	// padding for ASCII art to align with profile data
	asciiArtLines := strings.Split(asciiArt, "\n")
	maxLength := 0
	for _, line := range asciiArtLines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Print ASCII art and profile data side by side
	for i := 0; i < len(asciiArtLines) || i < len(data); i++ {
		asciiLine := ""
		if i < len(asciiArtLines) {
			asciiLine = asciiArtLines[i]
		}
		dataLine := ""
		if i < len(data) {
			dataLine = data[i]
		}

		// Align ASCII art and profile data vertically
		fmt.Printf("%-40s %s\n", asciiLine, dataLine)
	}
}
