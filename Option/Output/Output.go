package OUTPUT

import (
	"Ascii-Art/BANNER"
	"errors"
	"flag"
	"os"
)

func OUTPUT() error {
	saisie := BANNER.Saisie()
	output := errors.New("error")
	h := flag.String("output", "text", "collect data")
	flag.Parse()
	if os.Args[3] == "standard" {
		output = os.WriteFile(*h, []byte((BANNER.Standard(saisie))), 0777)
	}
	if os.Args[3] == "shadow" {
		output = os.WriteFile(*h, []byte((BANNER.Shadow(saisie))), 0777)
	}
	if os.Args[3] == "thinkertoy" {
		output = os.WriteFile(*h, []byte((BANNER.Thinkertoy(saisie))), 0777)
	}
	if output != nil {
		return errors.New("error")
	}
	return output
}
