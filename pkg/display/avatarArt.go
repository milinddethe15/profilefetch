package display

import (
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func avatarArt(url string) string {
	flags := aic_package.Flags{
		CharBackgroundColor: true,
		Colored:             true,
		CustomMap:           " .:-=+*%@#",
	}
	art, err := aic_package.Convert(url, flags)
	if err != nil {
		return ""
	}
	return art
}
