package Display

import ("fmt"
"os"
"log"
"strings")

func Display (line string) string{
	text := ""
	saisie := os.Args[2]
	Tab1 := strings.Split(saisie, "\\n")
	d := strings.Count(saisie, "\\n")
	if os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy" {
		if (len(os.Args[1:]) == 3) && os.Args[3] == "standard" {
			File, err := os.ReadFile("Files/standard.txt")
			Tab := strings.Split(string(File), "\n")
			if err != nil {
				return "\033[0mError: fichier standard.txt introuvable\n "
			}
			for index := 0; index <= d; index++ {
				for i := 0; i < 8; i++ {
					for _, word := range Tab1[index] {
						if word > 127 {
							if i > 6 {
								log.Fatal("\rerror :  verifiez votre argument (", os.Args[2], ") \"", string(word), "\" n'appartient pas aux caracteres man ascii attendue")
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
				if Tab1[index] == "" && index == 0 && float32(len(saisie))/2 != float32(d) {
					text = "\n"
				}
			}
		}
	}else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --color=red \"text\" standard \nBANNER List: \n-standard\n-shadow\n-thinkertoy")
		os.Exit(1)
	}
	return text
}