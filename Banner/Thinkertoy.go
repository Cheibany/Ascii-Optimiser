package Banner

import (
	"log"
	"os"
	"strings"
)

func Thinkertoy(Input string) string {
	text := ""
	Tab1 := strings.Split(Input, "\\n")
	d := strings.Count(Input, "\\n")
	File, err := os.ReadFile("Files/thinkertoy.txt")
	Tab := strings.Split(string(File), "\r\n")
	if err != nil {
		return "Error: fichier thinkertoy.txt introuvable\n"
	}
	for index := 0; index <= d; index++ {
		for i := 0; i < 8; i++ {
			for _, word := range Tab1[index] {
				if word > 127 {
					if i > 6 {
						log.Fatal("\rerror :  verifiez votre argument (", os.Args[1], ") \"", string(word), "\" n'appartient pas aux caracteres man ascii attendue\n")
					}
				} else {
					text += Tab[(((word-31)*9)-8)+rune(i)]
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
