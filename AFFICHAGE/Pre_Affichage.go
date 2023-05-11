package Affichage

import (
	"Ascii-Art/Banner"
	"Ascii-Art/Option/Color"
	"Ascii-Art/Option/Justify"
	Option "Ascii-Art/Option/Output"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Affichage() string {
	text := ""
	if len(os.Args[1:]) < 1 || len(os.Args[1:]) > 3 {
		return "\n\t\t\033[31m~FS~\033[0m\n\n\033[32musage:\033[0m go run . [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . \"HELLO\" shadow\n\n\t\t\033[31m~COLOR~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING]\n\n\033[32mexample:\033[0m go run . --color=red \"HELLO\"\n\n\t\t\033[31m~OUTPUT~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . --output=filename.txt \"HELLO\" thinkertoy\n\n"
	}
	if len(os.Args[1:]) == 1 {
		text = Banner.Standard(Banner.Saisie())
	}
	if len(os.Args[1:]) == 2 {
		saisie := Banner.Saisie()
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
			return "\n\t\t\033[31m~FS~\033[0m\n\n\033[32musage:\033[0m go run . [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . \"HELLO\" shadow\n\n\t\t\033[31m~COLOR~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING]\n\n\033[32mexample:\033[0m go run . --color=red \"HELLO\"\n\n\t\t\033[31m~OUTPUT~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . --output=filename.txt \"HELLO\" thinkertoy\n\n"
		}
	}
	if len(os.Args[1:]) == 3 {
		if strings.HasPrefix(os.Args[1], "--output=") || strings.HasPrefix(os.Args[1], "--color=") || strings.HasPrefix(os.Args[1], "--align=") {
			if strings.HasPrefix(os.Args[1], "--output=") {
				if strings.HasPrefix(os.Args[1], "--output=") {
					if Option.OUTPUT() != nil {
						return fmt.Sprint("\n❌Invalid Banner : ", os.Args[3], "\n\n\t\033[32mBanner List\033[0m\n\n-standard\n\n-shadow\n\n-thinkertoy\n\n")
					}
				}
			}
			if strings.HasPrefix(os.Args[1], "--color=") {
				if Color.Color() == "Invalid Color" {
					h := flag.String("color", "text", "here")
					flag.Parse()
					fmt.Print("❌Invalid Color: ", *h, "\n\n\t\tColor List\033[0m\n-black\n-red\n-green\n-yellow\n-blue\n-violet\n-cyan\n-white\n-orange\n-pink\n-marron\n")
					os.Exit(1)
				}
				return Banner.Standard(Banner.Saisie())
			}
			if strings.HasPrefix(os.Args[1], "--align=") {
				Justify.Justify()
			}
		} else {
			return "\n\t\t\033[31m~FS~\033[0m\n\n\033[32musage:\033[0m go run . [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . \"HELLO\" shadow\n\n\t\t\033[31m~COLOR~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING]\n\n\033[32mexample:\033[0m go run . --color=red \"HELLO\"\n\n\t\t\033[31m~OUTPUT~\033[0m\n\n\033[32musage:\033[0m go run . [OPTION] [STRING] [BANNER]\n\n\033[32mexample:\033[0m go run . --output=filename.txt \"HELLO\" thinkertoy\n\n"
		}
	}
	return text
}
