package display

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v60/github"
)

func DisplayOutput(username string, user *github.User) {
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

	
	profileData := []string{
		title(username + "@github"),
		strings.Repeat("-", 30),
		fmt.Sprintf("%s%s", label("Name: "), info(user.GetName())),
		fmt.Sprintf("%s%s", label("Email: "), info(user.GetEmail())),
		fmt.Sprintf("%s%s", label("Bio: "), info(user.GetBio())),
		fmt.Sprintf("%s%s", label("Followers: "), info(fmt.Sprintf("%d", user.GetFollowers()))),
		fmt.Sprintf("%s%s", label("Following: "), info(fmt.Sprintf("%d", user.GetFollowing()))),
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
	for i := 0; i < len(asciiArtLines) || i < len(profileData); i++ {
		asciiLine := ""
		if i < len(asciiArtLines) {
			asciiLine = asciiArtLines[i]
		}
		profileLine := ""
		if i < len(profileData) {
			profileLine = profileData[i]
		}

		// Align ASCII art and profile data vertically
		fmt.Printf("%-40s %s\n", asciiLine, profileLine)
	}
}
