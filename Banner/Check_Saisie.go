package Banner

import (
	"os"
	"strings"
)

func Saisie() string {
	var saisie string
	if len(os.Args[1:]) == 2 {
		saisie = os.Args[1]
	}
	if len(os.Args[1:]) == 3 || strings.HasPrefix(os.Args[1], "--color=") {
		saisie = os.Args[2]
	}
	if len(os.Args[1:]) == 3 && strings.HasPrefix(os.Args[1], "--color=") {
		saisie = os.Args[3]
	}
	return saisie
}
