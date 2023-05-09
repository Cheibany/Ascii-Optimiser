package Justify

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var alignType string

func init() {
	flag.StringVar(&alignType, "align", "left", "Type d'alignement (center, left, right, justify)")
	flag.Parse()
}

func Justify() {
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("Exemple: go run . --align=right quelque chose de standard")
		os.Exit(1)
	}
	text := args[0]

	// Lire la largeur de la console
	consoleWidth, err := getConsoleWidth()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch alignType {
	case "centre":
		fmt.Println(centreAlign(text, consoleWidth))
	case "gauche":
		fmt.Println(leftAlign(text))
	case "droite":
		fmt.Println(rightAlign(text, consoleWidth))
	case "justifier":
		fmt.Println(justifyAlign(text, consoleWidth))
	default:
		fmt.Println("Type d'alignement invalide!")
		os.Exit(1)
	}
}

// Obtenir la largeur de la console
func getConsoleWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	size := strings.Split(strings.TrimSpace(string(out)), " ")
	cols := size[1]
	w, _ := strconv.Atoi(cols)

	return w, nil
}

// Alignement à gauche
func leftAlign(text string) string {
	return text
}

// Alignement à droite
func rightAlign(text string, consoleWidth int) string {
	spaces := consoleWidth - len(text)
	if spaces <= 0 {
		return text
	}
	return fmt.Sprintf("%s%s", strings.Repeat(" ", spaces), text)
}

// Alignement centré
func centreAlign(text string, consoleWidth int) string {
	width := len(text)
	spaces := consoleWidth - width
	if spaces <= 0 {
		return text
	}
	left := spaces / 2
	right := spaces - left
	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", left), text, strings.Repeat(" ", right))
}

// Justifier l'alignement
func justifyAlign(text string, consoleWidth int) string {
	words := strings.Fields(text)
	numWords := len(words)
	if numWords == 1 {
		return leftAlign(text)
	}
	spacesNeeded := consoleWidth - len(text)
	spacesBetweenWords := spacesNeeded / (numWords - 1)
	extraSpaces := spacesNeeded % (numWords - 1)
	var justifiedText strings.Builder
	for i, word := range words {
		if i > 0 {
			spaces := spacesBetweenWords
			if extraSpaces > 0 {
				spaces++
				extraSpaces--
			}
			justifiedText.WriteString(strings.Repeat(" ", spaces))
		}
		justifiedText.WriteString(word)
	}
	return justifiedText.String()
}
