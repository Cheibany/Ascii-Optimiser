package Color

import (
	"os"
	"strings"
)

func Color() string {
	choix := ""
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
	// choix := flag.String("color", "text", "letter to color")
	// flag.Parse()
	// switch *choix {
	// case "black", "#000000", "r:0,g:0,b:0":
	// 	{
	// 		black := "\033[30m"
	// 		choix = &black
	// 	}
	// case "red", "#FF0000", "r:255,g:0,b:0":
	// 	{
	// 		red := "\033[31m"
	// 		choix = &red
	// 	}
	// case "green", "#008000", "r:0,g:128,b:0":
	// 	{
	// 		green := "\033[32m"
	// 		choix = &green
	// 	}
	// case "yellow", "#FFFF00", "r:255,g:255,b:0":
	// 	{
	// 		yellow := "\033[33m"
	// 		choix = &yellow
	// 	}
	// case "blue", "#0000FF", "r:0,g:0,b:255":
	// 	{
	// 		blue := "\033[34m"
	// 		choix = &blue
	// 	}
	// case "violet", "#800080", "r:128,g:0,b:128":
	// 	{
	// 		violet := "\033[35m"
	// 		choix = &violet
	// 	}
	// case "cyan", "#00FFFF":
	// 	{
	// 		cyan := "\033[36m"
	// 		choix = &cyan
	// 	}
	// case "white", "#FFFFFF", "r:255,g:255,b:255":
	// 	{
	// 		white := "\033[37m"
	// 		choix = &white
	// 	}
	// case "orange", "FFA500", "r:255,g:165,b:0":
	// 	{
	// 		orange := "\033[38;2;255;165;0m"
	// 		choix = &orange
	// 	}
	// case "marron", "#800000", "r:128,g:0,b:0":
	// 	{
	// 		marron := "\033[38;2;128;0;0m"
	// 		choix = &marron
	// 	}
	// case "pinks", "#FFB6C1", "r:,g:,b:":
	// 	{
	// 		pink := "\033[38;2;255;182;193m"
	// 		choix = &pink
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Color List:\n1-nom: black\thex: #000000\trgb: r:0,g:0,b:0 \n2-nom: red\thex: #FF0000\trgb: r:255,g:0,b:0\n3-green\n4-yellow\n5-blue\n6-violet\n7-cyan\n8-orange\n9-marron\n10-pink")
	// 		os.Exit(1)
	// 	}
	// }
	// return *choix
}
