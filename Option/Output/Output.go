package Output

import (
	"Ascii-Art/Banner"
	"errors"
	"flag"
	"os"
)

func OUTPUT() error {
	saisie := Banner.Saisie()
	output := errors.New("error")
	h := flag.String("output", "text", "collect data")
	flag.Parse()
	if os.Args[3] == "standard" {
		output = os.WriteFile(*h, []byte((Banner.Standard(saisie))), 0777)
	}
	if os.Args[3] == "shadow" {
		output = os.WriteFile(*h, []byte((Banner.Shadow(saisie))), 0777)
	}
	if os.Args[3] == "thinkertoy" {
		output = os.WriteFile(*h, []byte((Banner.Thinkertoy(saisie))), 0777)
	}
	if output != nil {
		return errors.New("error")
	}
	return output
}
