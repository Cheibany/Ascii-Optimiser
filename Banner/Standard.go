package BANNER

import (
	"Ascii-Art/OPTION/COLOR"
	"log"
	"os"
	"strings"
)

func Standard(Input string) string {
	text := ""
	Tab1 := strings.Split(Input, "\\n")
	d := strings.Count(Input, "\\n")
	File, err := os.ReadFile("Files/standard.txt")
	Tab := strings.Split(string(File), "\n")
	if err != nil {
		return "Error: fichier standard.txt introuvable\n"
	}
	for index := 0; index <= d; index++ {
		for i := 0; i < 8; i++ {
			for _, word := range Tab1[index] {
				if word > 127 {
					if i > 6 {
						log.Fatal("\rerror :  verifiez votre argument (", Input, ") \"", string(word), "\" n'appartient pas aux caracteres man ascii attendue")
					}
				} else {
					if strings.HasPrefix(os.Args[1], "--color=") && len(os.Args[1:]) == 3 {
						if strings.Contains(os.Args[2], string(word)) {
							text += COLOR.Color() + Tab[(((word-31)*9)-8)+rune(i)] + "\033[0m"
						} else {
							text += Tab[(((word-31)*9)-8)+rune(i)]
						}
					} else {
						text += Tab[(((word-31)*9)-8)+rune(i)]
					}
				}
			}
			if Tab1[index] != "" {
				text += "\n"
			}
		}
		if Tab1[index] == "" && index != 0 {
			text += "\n"
		}
		if Tab1[index] == "" && index == 0 && float32(len(Input))/2 != float32(d) {
			text = "\n"
		}
	}
	return text
}
