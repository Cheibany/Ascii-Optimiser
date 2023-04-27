package AFFICHAGE

import (
	"Ascii-Art/Banner"
	"Ascii-Art/Option/Color"
	Option "Ascii-Art/Option/Output"
	"fmt"
	"os"
	"strings"
)

func Affichage() string {
	text := ""
	saisie := Banner.Saisie()
	if len(os.Args[1:]) == 2 {
		if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" || strings.HasPrefix(os.Args[1], "--color=") {
			if os.Args[2] == "standard" {
				text = Banner.Standard(saisie)
			}
			if os.Args[2] == "shadow" {
				text = Banner.Shadow(saisie)
			}
			if os.Args[2] == "thinkertoy" {
				text = Banner.Thinkertoy(saisie)
			}
			if strings.HasPrefix(os.Args[1], "--color=") {
				text = fmt.Sprint(Color.Color(), Banner.Standard(saisie), "\033[0m")
			}
		} else {
			if strings.Contains(os.Args[1], "color") {
				return "usage: go run . [OPTION] [STRING]\nexample: go run . --color=red \"HELLO\"\n"
			}
			if strings.Contains(os.Args[1], "output") {
				return "usage: go run . [OPTION] [STRING] [BANNER]\nexample: go run . --output=filename.txt \"HELLO\" standard\n"
			}
			return "\t\tBanner List:\n-standard\n-shadow\n-thinkertoy\n"

		}
	}
	if len(os.Args[1:]) == 3 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			if strings.HasPrefix(os.Args[1], "--output=") {
				if Option.Output() != nil {
					return "error output"
				}
			}
		}
		if strings.HasPrefix(os.Args[1], "--color=") {
			return Banner.Standard(Banner.Saisie())
		}
	}
	//else {
	// 	return "usage: go run . [STRING] [BANNER]\nexample: go run . \"HELLO\" standard\n"
	// }
	return text
}
