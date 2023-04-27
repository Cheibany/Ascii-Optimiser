package main

import (
	"fmt"
	"log"
	"os"
	"display/Display"
)

func main() {
	if len(os.Args[1:]) == 3 {
		saisie := os.Args[2]

		fmt.Print(Display.Display(saisie))
	} else {
		log.Print("\rerror: verifie le nombre d'argument saisie\n\tUsage: go run . [OPTION] [STRING] [BANNER]\n\tEXAMPLE: go run . --color=red your_text standard")
	}
}
