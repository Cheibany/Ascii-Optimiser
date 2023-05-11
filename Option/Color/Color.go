package Color

import (
	"os"
	"strings"
)

func Color() string {
	choix := "Invalid Color"
	if strings.HasPrefix(os.Args[1], "--color=") {
		if strings.HasSuffix(os.Args[1], "=black") {
			black := "\033[30m"
			choix = black
		}
		if strings.HasSuffix(os.Args[1], "=red") {
			red := "\033[31m"
			choix = red
		}
		if strings.HasSuffix(os.Args[1], "=green") {
			green := "\033[32m"
			choix = green
		}
		if strings.HasSuffix(os.Args[1], "=yellow") {
			yellow := "\033[33m"
			choix = yellow
		}
		if strings.HasSuffix(os.Args[1], "=blue") {
			blue := "\033[34m"
			choix = blue
		}
		if strings.HasSuffix(os.Args[1], "=violet") {
			violet := "\033[35m"
			choix = violet
		}
		if strings.HasSuffix(os.Args[1], "=cyan") {
			cyan := "\033[36m"
			choix = cyan
		}
		if strings.HasSuffix(os.Args[1], "=white") {
			white := "\033[37m"
			choix = white
		}
		if strings.HasSuffix(os.Args[1], "=orange") {
			orange := "\033[38;2;255;165;0m"
			choix = orange
		}
		if strings.HasSuffix(os.Args[1], "=pink") {
			pink := "\033[38;2;255;182;193m"
			choix = pink
		}
		if strings.HasSuffix(os.Args[1], "=marron") || strings.HasSuffix(os.Args[1], "=#800000") {
			marron := "\033[38;2;128;0;0m"
			choix = marron
		}
	}
	return choix
}
